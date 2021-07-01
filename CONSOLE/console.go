package CONSOLE

import (
	"6502/CORE"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/cassianoperin/pseudo-terminal-go/terminal"
)

type instructuction struct {
	code        byte
	bytes       byte
	cycles      byte
	description string
	memory_mode string
}

type breakpoint struct {
	location string
	value    uint16
}

var breakpoints []breakpoint
var opcode_map = []instructuction{
	{0x0A, 1, 2, "ASL", "accumulator"},
	{0x18, 1, 2, "CLC", "implied"},
	{0xD8, 1, 2, "CLD", "implied"},
	{0x58, 1, 2, "CLI", "implied"},
	{0xB8, 1, 2, "CLV", "implied"},
	{0xCA, 1, 2, "DEX", "implied"},
	{0x88, 1, 2, "DEY", "implied"},
	{0xE8, 1, 2, "INX", "implied"},
	{0xC8, 1, 2, "INY", "implied"},
	{0x4A, 1, 2, "LSR", "accumulator"},
	{0xEA, 1, 2, "NOP", "implied"},
	{0x2A, 1, 2, "ROL", "accumulator"},
	{0x38, 1, 2, "SEC", "implied"},
	{0xF8, 1, 2, "SED", "implied"},
	{0x78, 1, 2, "SEI", "implied"},
	{0xAA, 1, 2, "TAX", "implied"},
	{0xA8, 1, 2, "TAY", "implied"},
	{0xBA, 1, 2, "TSX", "implied"},
	{0x8A, 1, 2, "TXA", "implied"},
	{0x9A, 1, 2, "TXS", "implied"},
	{0x98, 1, 2, "TYA", "implied"},
	{0x69, 2, 2, "ADC", "immediate"},
	{0x65, 2, 3, "ADC", "zeropage"},
	{0x75, 2, 4, "ADC", "zeropage,X"},
	{0x6D, 3, 4, "ADC", "absolute"},
	{0x7D, 3, 4, "ADC", "absolute,X"},
	{0x79, 3, 4, "ADC", "absolute,Y"},
	{0x61, 2, 6, "ADC", "(indirect,X)"},
	{0x71, 2, 5, "ADC", "(indirect),Y"},
	{0x29, 2, 2, "AND", "immediate"},
	{0x25, 2, 3, "AND", "zeropage"},
	{0x35, 2, 4, "AND", "zeropage,X"},
	{0x2D, 3, 4, "AND", "absolute"},
	{0x3D, 3, 4, "AND", "absolute,X"},
	{0x39, 3, 4, "AND", "absolute,Y"},
	{0x21, 2, 6, "AND", "(indirect,X)"},
	{0x31, 2, 5, "AND", "(indirect),Y"},
	{0x2C, 3, 4, "BIT", "absolute"},
	{0x24, 2, 3, "BIT", "zeropage"},
	{0xC5, 2, 3, "CMP", "zeropage"},
	{0xC9, 2, 2, "CMP", "immediate"},
	{0xD5, 2, 4, "CMP", "zeropage,X"},
	{0xCD, 3, 4, "CMP", "absolute"},
	{0xD9, 3, 4, "CMP", "absolute,Y"},
	{0xDD, 3, 4, "CMP", "absolute,X"},
	{0xD1, 2, 5, "CMP", "(indirect),Y"},
	{0xC1, 2, 6, "CMP", "(indirect,X)"},
}

func StartConsole() {

	// Reset system
	CORE.Reset()

	term, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}
	defer term.ReleaseFromStdInOut() // defer this
	fmt.Printf("\n------------------------ Console mode ------------------------\n\nType \"Ctrl-Q\" to quit, \"help\" for more options\n\n")

	// Print Header
	printHeader()

	term.SetPrompt("# ")
	line, err := term.ReadLine()
	for {
		if err == io.EOF {
			term.Write([]byte(line))
			fmt.Println()
			return
		}
		if (err != nil && strings.Contains(err.Error(), "control-c break")) || len(line) == 0 {
			line, err = term.ReadLine()
		} else {

			// Command Interpreter
			CommandInterpreter(line)

			// term.Write([]byte(line + "\r\n"))
			line, err = term.ReadLine()
		}
	}
	term.Write([]byte(line))
}

// Print the debub information in console
func print_debug_console(opcode_map []instructuction) {
	for i := 0; i < len(opcode_map); i++ {

		if CORE.Memory[CORE.PC] == opcode_map[i].code {
			opc_string, opc_operand := CORE.Debug_decode_console(opcode_map[i].bytes)
			fmt.Printf("\n--> $%04X\t%s %s\t%s %s (%s)\n\n", CORE.PC, opc_string, opc_operand, opcode_map[i].description, opc_operand, opcode_map[i].memory_mode)
		}

	}
}

// Execute the necessary cycles for next instruction
func step(opcode_map []instructuction) {
	// Print the opcode debug
	print_debug_console(opcode_map)

	for !CORE.NewInstruction {
		CORE.CPU_Interpreter()
	}

	// Reset new instruction flag
	CORE.NewInstruction = false

	// Print the Header
	printHeader()
}

// Execute the necessary cycles for next instruction without print on console
func step_without_debug(opcode_map []instructuction) {

	for !CORE.NewInstruction {
		CORE.CPU_Interpreter()
	}

	// Reset new instruction flag
	CORE.NewInstruction = false

}

// Print Help Menu
func printHelp() {
	fmt.Printf("\n\tquit\t\t\t\t\t\tQuit console\n\thelp\t\t\t\t\t\tPrint help menu\n\tstep\t\t\t\t\t\tExecute current opcode\n\tadd_breakpoint <PC|A|X|Y|MEM|CYCLE>=<Value>\tAdd a breakpoint\n\tdel_breakpoint <index>\t\t\t\tDelete a breakpoint\n\tshow_breakpoints\t\t\t\tShow breakpoints\n\tstep <value>\t\t\t\t\tExecute <value> opcodes\n\trun\t\t\t\t\t\tRun the emulator\n\n")
}

func printAddBrkErr() {
	fmt.Println("Usage: add_breakpoint < PC | A | X | Y | MEM | CYCLE > = <Value>\n")
}

// Print current Processor Information
func printHeader() {
	fmt.Printf("  -------------------------------------------------------------------------\n")
	fmt.Printf("  |   PC\tA\tX\tY\tSP\tNV-BDIZC\tCycle\n")
	fmt.Printf("  |   %04X\t%02X\t%02X\t%02X\t%02X\t%d%d%d%d%d%d%d%d\t%d\n", CORE.PC, CORE.A, CORE.X, CORE.Y, CORE.SP, CORE.P[7], CORE.P[6], CORE.P[5], CORE.P[4], CORE.P[3], CORE.P[2], CORE.P[1], CORE.P[0], CORE.Cycle)
	fmt.Printf("  -------------------------------------------------------------------------\n\n")
}

// Remove a value from a slice
func RemoveIndex(s []breakpoint, index int) []breakpoint {
	return append(s[:index], s[index+1:]...)
}

// Interpreter
func CommandInterpreter(text string) {

	if strings.Compare("quit", text) == 0 || strings.Compare("Quit", text) == 0 || strings.Compare("exit", text) == 0 || strings.Compare("Exit", text) == 0 {
		fmt.Printf("Exiting console.\n")
		os.Exit(0)

	} else if strings.Compare("help", text) == 0 || strings.Compare("Help", text) == 0 { // Help
		printHelp()

	} else if strings.Compare("step", text) == 0 { // STEP

		// Execute one instruction
		step(opcode_map)

	} else if strings.HasPrefix(text, "step ") { // STEP WITH ARGS

		// Convert string value to integer
		value, err := strconv.Atoi(strings.TrimPrefix(text, "step "))
		if err != nil {
			// handle error
			fmt.Printf("Invalid value: %s\n", strings.TrimPrefix(text, "step "))
		} else {

			// Number os steps
			for i := 0; i < value; i++ {

				// Execute one instruction
				step(opcode_map)

			}

		}
	} else if strings.HasPrefix(text, "add_breakpoint") { // ADD BREAKPOINT

		var tmp_string, tmp_string2 []string

		tmp_string = strings.Split(text, " ") // First split command and values

		// Test the command syntax
		if len(tmp_string) == 1 || len(tmp_string) > 2 {

			// Print add_breakpoint usage
			printAddBrkErr()

		} else {

			// If command syntax is ok, test the parameter syntax
			tmp_string2 = strings.Split(tmp_string[1], "=") // After, split the argument in LOCATION=VALUE
			if len(tmp_string2) == 1 || len(tmp_string2) > 2 || tmp_string2[1] == "" || tmp_string2[0] == "" {

				// Print add_breakpoint usage
				printAddBrkErr()

			} else {

				location := strings.ToUpper(tmp_string2[0])

				// Validate the value of locations
				if location == "PC" || location == "A" || location == "X" || location == "Y" || location == "MEM" {

					value, err := strconv.Atoi(tmp_string2[1])

					if err != nil {
						fmt.Println("Invalid address.")
					} else {
						if value <= 65535 && value >= 0 {
							breakpoints = append(breakpoints, breakpoint{strings.ToUpper(tmp_string2[0]), uint16(value)})
							fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
						} else {
							fmt.Println("Invalid address.")
						}
					}
				} else {

					// Print add_breakpoint usage
					printAddBrkErr()
				}

			}

		}

	} else if strings.HasPrefix(text, "del_breakpoint") { // DELETE BREAKPOINT

		if len(strings.Split(text, " ")) == 1 || len(strings.Split(text, " ")) > 2 {
			fmt.Println("Usage: del_breakpoint <breakpoint number>\n")
		} else {

			tmp_string := strings.Split(text, " ")

			value, err := strconv.Atoi(tmp_string[1])
			if err != nil {
				// handle error
				fmt.Printf("Invalid value %s\n\n", tmp_string[1])
			} else {
				if value < len(breakpoints) {
					breakpoints = RemoveIndex(breakpoints, value)
					fmt.Printf("Breakpoint %d removed.\n\n", value)
				} else {
					fmt.Printf("Breakpoint not found\n\n")
				}

			}
		}

	} else if strings.Compare("show_breakpoints", text) == 0 { // SHOW BREAKPOINTS

		for i := 0; i < len(breakpoints); i++ {
			fmt.Printf("Breakpoint %d: %s=%d\n", i, breakpoints[i].location, breakpoints[i].value)
			if i == len(breakpoints)-1 {
				fmt.Println()
			}
		}

		if len(breakpoints) == 0 {
			fmt.Printf("No Breakpoint found.\n\n")
		}

	} else if strings.Compare("run", text) == 0 { // RUN

		var (
			current_PC uint16
			loop_count uint16 = 0
		)

		for loop_count < CORE.Loop_detection { // Add breakpoint control here

			// Keep track of current PC
			current_PC = CORE.PC

			select {
			case <-CORE.Second_timer: // Show the header and debug each second

				// Execute one instruction
				step(opcode_map)

			default: // Just run the CPU

				// Execute one instruction without print
				step_without_debug(opcode_map)

			}

			// Increase the loop counterß
			if current_PC == CORE.PC {
				loop_count++
			}

		}

		fmt.Printf("Loop detected on PC=%04X (%d repetitions)\n", CORE.PC, CORE.Loop_detection)

		// Print Header
		printHeader()

	} else { // Command not found
		fmt.Printf("Command not found\n\n")
	}

}
