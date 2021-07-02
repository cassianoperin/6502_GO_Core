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

var (
	step_limit  int = 1000
	run_limit   int = 100000
	breakpoints []breakpoint
	opcode_map  = []instructuction{
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
)

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
	fmt.Printf("\n\tquit\t\t\t\t\t\tQuit console\n\thelp\t\t\t\t\t\tPrint help menu\n\tstep\t\t\t\t\t\tExecute current opcode\n\tstep_limit <value>\t\t\t\tDefine the maximum steps allowed\n\tadd_breakpoint <PC|A|X|Y|CYCLE>=<Value>\tAdd a breakpoint\n\tdel_breakpoint <index>\t\t\t\tDelete a breakpoint\n\tshow_breakpoints\t\t\t\tShow breakpoints\n\tstep <value>\t\t\t\t\tExecute <value> opcodes\n\trun\t\t\t\t\t\tRun the emulator\n\trun_limit <value>\t\t\t\tDefine the maximum steps allowed in RUN\n\n")
}

func printAddBrkErr() {
	fmt.Printf("Usage: add_breakpoint < PC | A | X | Y | CYCLE > = <Value>\n\n")
}

func printStepLimitErr() {
	fmt.Printf("Usage: step_limit <Value>\n\n")
}

func printRunLimitErr() {
	fmt.Printf("Usage: run_limit <Value>\n\n")
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

	// Remove duplicated spaces
	text = strings.Join(strings.Fields(strings.TrimSpace(text)), " ")

	if strings.Contains(text, "quit") || strings.Contains(text, "Quit") || strings.Contains(text, "exit") || strings.Contains(text, "Exit") {
		fmt.Printf("Exiting console.\n")
		os.Exit(0)

	} else if strings.Contains(text, "help") || strings.Contains(text, "Help") { // Help
		printHelp()

	} else if strings.Contains(text, "step") { // STEP

		if strings.Compare("step", text) == 0 { // STEP

			// Execute one instruction
			step(opcode_map)

		} else if strings.HasPrefix(text, "step_limit") {

			tmp_string := strings.Split(text, " ")

			// Test the command syntax
			if len(tmp_string) == 1 {

				// Show current value
				fmt.Printf("Current Step Limit = %d\n\n", step_limit)

			} else if len(tmp_string) > 2 {

				// Print step_limit usage
				printStepLimitErr()

			} else {

				// Convert string value to integer
				value, err := strconv.Atoi(strings.TrimPrefix(text, "step_limit "))
				if err != nil {
					// handle error
					fmt.Printf("Invalid value: %s\n\n", strings.TrimPrefix(text, "step_limit "))
				} else {
					step_limit = value
					fmt.Printf("New step limit = %d\n\n", step_limit)
				}

			}

		} else {

			var breakpoint_flag bool

			// Convert string value to integer
			value, err := strconv.Atoi(strings.TrimPrefix(text, "step "))
			if err != nil {
				// handle error
				fmt.Printf("Invalid value: %s\n\n", strings.TrimPrefix(text, "step "))
			} else {

				// Number os steps
				if value <= step_limit {
					for i := 0; i < value; i++ {

						// Execute one instruction
						step(opcode_map)

						// Check Breakpoints
						breakpoint_flag = check_breakpoints(breakpoint_flag)

						// Exit for loop if breakpoint has been found
						if breakpoint_flag {
							break
						}

					}
				} else {
					fmt.Printf("Current step limit = %d\n\n", step_limit)
				}

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
				if location == "PC" || location == "A" || location == "X" || location == "Y" || location == "CYCLE" {

					value, err := strconv.Atoi(tmp_string2[1])

					if err != nil {
						fmt.Println("Invalid value.")
					} else {
						// Value limits
						if location == "PC" {
							if value <= 65535 && value >= 0 {
								breakpoints = append(breakpoints, breakpoint{strings.ToUpper(tmp_string2[0]), uint16(value)})
								fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
							} else {
								fmt.Println("Invalid value.")
							}
						}

						if location == "A" || location == "X" || location == "Y" {
							if value <= 255 && value >= 0 {
								breakpoints = append(breakpoints, breakpoint{strings.ToUpper(tmp_string2[0]), uint16(value)})
								fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
							} else {
								fmt.Println("Invalid value.")
							}
						}

						if location == "CYCLE" {
							if value >= 0 {
								breakpoints = append(breakpoints, breakpoint{strings.ToUpper(tmp_string2[0]), uint16(value)})
								fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
							} else {
								fmt.Println("Invalid value.")
							}
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

	} else if strings.Contains(text, "show_breakpoints") { // SHOW BREAKPOINTS

		tmp_string := strings.Split(text, " ")

		// Test the command syntax
		if len(tmp_string) == 1 {

			for i := 0; i < len(breakpoints); i++ {
				fmt.Printf("Breakpoint %d: %s=%d\n", i, breakpoints[i].location, breakpoints[i].value)
				if i == len(breakpoints)-1 {
					fmt.Println()
				}
			}

			if len(breakpoints) == 0 {
				fmt.Printf("No Breakpoint found.\n\n")
			}

		} else {

			// Print add_breakpoint usage
			fmt.Println("show_breakpoints doesn't accept arguments\n")

		}

	} else if strings.Contains(text, "run") { // RUN

		if strings.Compare("run", text) == 0 { // RUN

			var (
				current_PC      uint16
				step_count      int    = 0
				loop_count      uint16 = 0
				breakpoint_flag bool
			)

			for loop_count < CORE.Loop_detection { // Add breakpoint control here

				// -------------------------- Start Checks --------------------------- //

				// Check Run step limits
				if step_count > run_limit {
					break // Exit for loop
				}

				// Check Breakpoints
				breakpoint_flag = check_breakpoints(breakpoint_flag)

				// Exit for loop if breakpoint has been found
				if breakpoint_flag {
					break
				}

				// -------------- Finish checks and return to execution -------------- //
				current_PC = CORE.PC

				select {
				case <-CORE.Second_timer: // Show the header and debug each second

					// Execute one instruction
					step(opcode_map)

				default: // Just run the CPU

					// Execute one instruction without print
					step_without_debug(opcode_map)

				}

				// Increase steps count
				step_count++

				// Check for run_limit and print debug message prior to quit loop
				if step_count > run_limit { // Print limit reached message
					fmt.Printf("Run limit reached (%d)\n\n", run_limit)
				}

				// Increase the loop counter
				if current_PC == CORE.PC {
					loop_count++
				}

				// Check for loops and print debug message prior to quit loop
				if loop_count >= CORE.Loop_detection {
					fmt.Printf("Loop detected on PC=%04X (%d repetitions)\n", CORE.PC, CORE.Loop_detection)
				}

			}

			// Print Header
			printHeader()

		} else if strings.HasPrefix(text, "run_limit") {

			tmp_string := strings.Split(text, " ")

			// Test the command syntax
			if len(tmp_string) == 1 {

				// Show current value
				fmt.Printf("Current Run Limit = %d\n\n", run_limit)

			} else if len(tmp_string) > 2 {

				// Print run_limit usage
				printRunLimitErr()

			} else {

				// Convert string value to integer
				value, err := strconv.Atoi(strings.TrimPrefix(text, "run_limit "))
				if err != nil {
					// handle error
					fmt.Printf("Invalid value: %s\n\n", strings.TrimPrefix(text, "run_limit "))
				} else {
					run_limit = value
					fmt.Printf("New run limit = %d\n\n", run_limit)
				}

			}

		} else { // Command not found
			fmt.Printf("Command not found\n\n")
		}

	} else { // Command not found
		fmt.Printf("Command not found\n\n")
	}

}

func check_breakpoints(break_flag bool) bool {
	// Check Breakpoints
	if len(breakpoints) > 0 {
		for i := 0; i < len(breakpoints); i++ {

			// ------ PC ------ //
			if breakpoints[i].location == "PC" {
				if CORE.PC == breakpoints[i].value {
					fmt.Printf("Breakpoint %d reached: %s=%d (0x%04X)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ A ------- //
			} else if breakpoints[i].location == "A" {
				if CORE.A == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=%d (0x%02X)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ X ------- //
			} else if breakpoints[i].location == "X" {
				if CORE.X == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=%d (0x%02X)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ Y ------- //
			} else if breakpoints[i].location == "Y" {
				if CORE.Y == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=%d (0x%02X)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ Cycle ------- //
			} else if breakpoints[i].location == "CYCLE" {
				if CORE.Cycle == uint64(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=%d (0x%02X)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}
			}

		}

	}

	return break_flag
}
