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
		{0xE0, 2, 2, "CPX", "immediate"},
		{0xE4, 2, 3, "CPX", "zeropage"},
		{0xEC, 3, 4, "CPX", "absolute"},
		{0xC0, 2, 2, "CPY", "immediate"},
		{0xC4, 2, 3, "CPY", "zeropage"},
		{0xCC, 3, 4, "CPY", "absolute"},
		{0x49, 2, 2, "EOR", "immediate"},
		{0x45, 2, 3, "EOR", "zeropage"},
		{0x55, 2, 4, "EOR", "zeropage,X"},
		{0x4D, 3, 4, "EOR", "absolute"},
		{0x5D, 3, 4, "EOR", "absolute,X"},
		{0x59, 3, 4, "EOR", "absolute,Y"},
		{0x41, 2, 6, "EOR", "(indirect,X)"},
		{0x51, 2, 5, "EOR", "(indirect),Y"},
		{0xA9, 2, 2, "LDA", "immediate"},
		{0xA5, 2, 3, "LDA", "zeropage"},
		{0xB9, 3, 4, "LDA", "absolute,Y"},
		{0xBD, 3, 4, "LDA", "absolute,X"},
		{0xB1, 2, 5, "LDA", "(indirect),Y"},
		{0xB5, 2, 4, "LDA", "zeropage,X"},
		{0xAD, 3, 4, "LDA", "absolute"},
		{0xA1, 2, 6, "LDA", "(indirect,X)"},
		{0xA2, 2, 2, "LDX", "immediate"},
		{0xA6, 2, 3, "LDX", "zeropage"},
		{0xB6, 2, 4, "LDX", "zeropage,Y"},
		{0xBE, 3, 4, "LDX", "absolute,Y"},
		{0xAE, 3, 4, "LDX", "absolute"},
		{0xA0, 2, 2, "LDY", "immediate"},
		{0xA4, 2, 3, "LDY", "zeropage"},
		{0xB4, 2, 4, "LDY", "zeropage,X"},
		{0xAC, 3, 4, "LDY", "absolute"},
		{0xBC, 3, 4, "LDY", "absolute,X"},
		{0x09, 2, 2, "ORA", "immediate"},
		{0x05, 2, 3, "ORA", "zeropage"},
		{0x15, 2, 4, "ORA", "zeropage,X"},
		{0x0D, 3, 4, "ORA", "absolute"},
		{0x1D, 3, 4, "ORA", "absolute,X"},
		{0x19, 3, 4, "ORA", "absolute,Y"},
		{0x01, 2, 6, "ORA", "(indirect,X)"},
		{0x11, 2, 5, "ORA", "(indirect),Y"},
		{0xE9, 2, 2, "SBC", "immediate"},
		{0xE5, 2, 3, "SBC", "zeropage"},
		{0xF5, 2, 4, "SBC", "zeropage,X"},
		{0xED, 3, 4, "SBC", "absolute"},
		{0xFD, 3, 4, "SBC", "absolute,X"},
		{0xF9, 3, 4, "SBC", "absolute,Y"},
		{0xE1, 2, 6, "SBC", "(indirect,X)"},
		{0xF1, 2, 5, "SBC", "(indirect),Y"},
		{0x95, 2, 4, "STA", "zeropage,X"},
		{0x85, 2, 3, "STA", "zeropage"},
		{0x99, 3, 5, "STA", "absolute,Y"},
		{0x8D, 3, 4, "STA", "absolute"},
		{0x91, 2, 6, "STA", "(indirect),Y"},
		{0x9D, 3, 5, "STA", "absolute,X"},
		{0x81, 2, 6, "STA", "(indirect,X)"},
		{0x86, 2, 3, "STX", "zeropage"},
		{0x96, 2, 4, "STX", "zeropage,Y"},
		{0x8E, 3, 4, "STX", "absolute"},
		{0x84, 2, 3, "STY", "zeropage"},
		{0x94, 2, 4, "STY", "zeropage,X"},
		{0x8C, 3, 4, "STY", "absolute"},
		{0x06, 2, 5, "ASL", "zeropage"},
		{0x16, 2, 6, "ASL", "zeropage,X"},
		{0x0E, 3, 6, "ASL", "absolute"},
		{0x1E, 3, 7, "ASL", "absolute,X"},
		{0xC6, 2, 5, "DEC", "zeropage"},
		{0xD6, 2, 6, "DEC", "zeropage,X"},
		{0xCE, 3, 6, "DEC", "absolute"},
		{0xDE, 3, 7, "DEC", "absolute,X"},
		{0xE6, 2, 5, "INC", "zeropage"},
		{0xF6, 2, 6, "INC", "zeropage,X"},
		{0xEE, 3, 6, "INC", "absolute"},
		{0xFE, 3, 7, "INC", "absolute,X"},
		{0x46, 2, 5, "LSR", "zeropage"},
		{0x56, 2, 6, "LSR", "zeropage,X"},
		{0x4E, 3, 6, "LSR", "absolute"},
		{0x5E, 3, 7, "LSR", "absolute,X"},
		{0x26, 2, 5, "ROL", "zeropage"},
		{0x36, 2, 6, "ROL", "zeropage,X"},
		{0x2E, 3, 6, "ROL", "absolute"},
		{0x3E, 3, 7, "ROL", "absolute,X"},
		{0x6A, 1, 2, "ROR", "accumulator"},
		{0x66, 2, 5, "ROR", "zeropage"},
		{0x76, 2, 6, "ROR", "zeropage,X"},
		{0x6E, 3, 6, "ROR", "absolute"},
		{0x7E, 3, 7, "ROR", "absolute,X"},
		{0x48, 1, 3, "PHA", "implied"},
		{0x08, 1, 3, "PHP", "implied"},
		{0x68, 1, 4, "PLA", "implied"},
		{0x28, 1, 4, "PLP", "implied"},
		{0x4C, 3, 3, "JMP", "absolute"},
		{0x6C, 3, 5, "JMP", "indirect"},
		{0x20, 3, 6, "JSR", "absolute"},
		{0x40, 1, 6, "RTI", "implied"},
		{0x60, 1, 6, "RTS", "implied"},
		{0x00, 1, 7, "BRK", "implied"},
		{0xD0, 2, 2, "BNE", "relative"},
		{0xF0, 2, 2, "BEQ", "relative"},
		{0x10, 2, 2, "BPL", "relative"},
		{0x30, 2, 2, "BMI", "relative"},
		{0x70, 2, 2, "BVS", "relative"},
		{0x50, 2, 2, "BVC", "relative"},
		{0xB0, 2, 2, "BCS", "relative"},
		{0x90, 2, 2, "BCC", "relative"},
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

	// term.SetPrompt("# ")
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
	// fmt.Printf("\n\tquit\t\t\t\t\t\tQuit console\n\thelp\t\t\t\t\t\tPrint help menu\n\t--\n\tstep\t\t\t\t\t\tExecute current opcode\n\tstep <value>\t\t\t\t\tExecute <value> opcodes\n\tstep_limit <value>\t\t\t\tDefine the maximum steps allowed\n\t--\n\tadd_breakpoint <PC|A|X|Y|CYCLE>=<Value>\t\tAdd a breakpoint\n\tdel_breakpoint <index>\t\t\t\tDelete a breakpoint\n\tshow_breakpoints\t\t\t\tShow breakpoints\n\t--\n\tmem || mem <address> || mem <address> <address>\tDump address values\n\t--\n\trun\t\t\t\t\t\tRun the emulator\n\trun_limit <value>\t\t\t\tDefine the maximum steps allowed in RUN\n\n")
	fmt.Printf("\n\tquit\t\t\t\t\t\tQuit console\n\thelp\t\t\t\t\t\tPrint help menu\n\t-")
	fmt.Printf("\n\tstep\t\t\t\t\t\tExecute current opcode\n\tstep <value>\t\t\t\t\tExecute <value> opcodes\n\tstep_limit <value>\t\t\t\tDefine the maximum steps allowed\n\t-")
	fmt.Printf("\n\tadd_breakpoint <PC|A|X|Y|CYCLE>=<Value>\t\tAdd a breakpoint\n\tdel_breakpoint <index>\t\t\t\tDelete a breakpoint\n\tshow_breakpoints\t\t\t\tShow breakpoints\n\t-")
	fmt.Printf("\n\tmem\t\t\t\t\t\tDump full memory\n\tmem <address>\t\t\t\t\tDump memory address\n\tmem <start address> <end address>\t\tDump memory address range\n\t-")
	fmt.Printf("\n\trun\t\t\t\t\t\tRun the emulator\n\trun_limit <value>\t\t\t\tDefine the maximum steps allowed in RUN\n\n")
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

					// Test if the value start if 0x or 0X
					if strings.HasPrefix(tmp_string2[1], "0x") || strings.HasPrefix(tmp_string2[1], "0X") {
						// fmt.Println("seria HEX")

						var hexaString string = tmp_string2[1]
						numberStr := strings.Replace(hexaString, "0x", "", -1)
						numberStr = strings.Replace(numberStr, "0X", "", -1)

						value, err := strconv.ParseInt(numberStr, 16, 64)

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
						// fmt.Println("seria DEC")

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
				fmt.Printf("Breakpoint %d: %s=0x%02X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
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

	} else if strings.HasPrefix(text, "mem") { // MEMORY

		tmp_string := strings.Split(text, " ")

		// Check command "mem"
		if strings.Compare("mem", tmp_string[0]) == 0 {

			if len(tmp_string) == 1 { // Without arguments (show all memory)
				fmt.Printf("\t00 01 02 03 04 05 06 07 08 09 0A 0B 0C 0D 0E 0F\n")
				fmt.Printf("\t-----------------------------------------------")
				for i := 0; i < len(CORE.Memory); i++ {

					// Break lines
					if i%16 == 0 {
						fmt.Printf("\n%04X\t", i)
					}

					// Print memory
					fmt.Printf("%02X ", CORE.Memory[i])

				}
				fmt.Println()

			} else if len(tmp_string) == 2 { // Without ONE argument (show this memory value)

				var mem1 int

				// Test if the value start if 0x or 0X
				if strings.HasPrefix(tmp_string[1], "0x") || strings.HasPrefix(tmp_string[1], "0X") {

					// HEXADECIMAL Input

					var hexaString string = tmp_string[1]
					numberStr := strings.Replace(hexaString, "0x", "", -1)
					numberStr = strings.Replace(numberStr, "0X", "", -1)

					tmp_value, err := strconv.ParseInt(numberStr, 16, 64)

					if err != nil {
						fmt.Println("Invalid value.")
					} else {
						// Convert to decimal and set mem1 value
						mem1 = int(tmp_value)
					}

				} else {

					// DECIMAL Input

					tmp_value, err := strconv.Atoi(tmp_string[1])
					if err != nil {
						// handle error
						fmt.Printf("Invalid value %s\n\n", tmp_string[1])
					} else {
						// Set mem1 value
						mem1 = int(tmp_value)
					}
				}

				// Print Memory Value
				if mem1 < 0 || mem1 >= len(CORE.Memory) {
					fmt.Printf("Invalid Address %d\n\n", mem1)
				} else {
					fmt.Printf("%02X\n\n", CORE.Memory[mem1])

				}

			} else if len(tmp_string) == 3 {

				var mem1, mem2 int
				var error_flag bool

				// Test if the FIRST ARGUMENT start if 0x or 0X
				if strings.HasPrefix(tmp_string[1], "0x") || strings.HasPrefix(tmp_string[1], "0X") {

					// FIRST ARGUMENT HEXADECIMAL Input

					var hexaString string = tmp_string[1]
					numberStr := strings.Replace(hexaString, "0x", "", -1)
					numberStr = strings.Replace(numberStr, "0X", "", -1)

					tmp_value, err := strconv.ParseInt(numberStr, 16, 64)

					if err != nil {
						fmt.Println("Invalid value.")
						error_flag = true
					} else {
						// Convert to decimal and set mem1 value
						mem1 = int(tmp_value)
					}

				} else {

					// FIRST ARGUMENT DECIMAL Input

					tmp_value, err := strconv.Atoi(tmp_string[1])
					if err != nil {
						// handle error
						fmt.Printf("Invalid value %s\n\n", tmp_string[1])
						error_flag = true
					} else {
						// Set mem1 value
						mem1 = int(tmp_value)
					}
				}

				// Test if the SECOND ARGUMENT start if 0x or 0X
				if strings.HasPrefix(tmp_string[2], "0x") || strings.HasPrefix(tmp_string[1], "0X") {

					// SECOND ARGUMENT HEXADECIMAL Input

					var hexaString string = tmp_string[2]
					numberStr := strings.Replace(hexaString, "0x", "", -1)
					numberStr = strings.Replace(numberStr, "0X", "", -1)

					tmp_value, err := strconv.ParseInt(numberStr, 16, 64)

					if err != nil {
						fmt.Println("Invalid value.")
						error_flag = true
					} else {
						// Convert to decimal and set mem1 value
						mem2 = int(tmp_value)
					}

				} else {

					// SECOND ARGUMENT DECIMAL Input

					tmp_value, err := strconv.Atoi(tmp_string[2])
					if err != nil {
						// handle error
						fmt.Printf("Invalid value %s\n\n", tmp_string[2])
						error_flag = true
					} else {
						// Set mem1 value
						mem2 = int(tmp_value)
					}
				}

				if !error_flag {
					// Print Memory Value
					if mem1 < 0 || mem1 >= len(CORE.Memory) {
						fmt.Printf("Invalid Address %d\n\n", mem1)
						error_flag = true
					} else if mem2 < 0 || mem2 >= len(CORE.Memory) {
						fmt.Printf("Invalid Address %d\n\n", mem2)
						error_flag = true
					} else if mem1 > mem2 {
						fmt.Printf("Start address should be less or equal end address\n\n")
						error_flag = true
					} else {
						for i := mem1; i <= mem2; i++ {
							fmt.Printf("%02X ", CORE.Memory[i])
						}
						fmt.Printf("\n\n")
					}
				}

			} else {
				fmt.Printf("Usage:\n   mem\n   mem <address>\n   mem <start address> <end address>\n\n")
			}
		} else {
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
					fmt.Printf("Breakpoint %d reached: %s=0x%04X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ A ------- //
			} else if breakpoints[i].location == "A" {
				if CORE.A == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=0x%02X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ X ------- //
			} else if breakpoints[i].location == "X" {
				if CORE.X == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=0x%02X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ Y ------- //
			} else if breakpoints[i].location == "Y" {
				if CORE.Y == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=0x%02X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ Cycle ------- //
			} else if breakpoints[i].location == "CYCLE" {
				if CORE.Cycle >= uint64(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=%d\t(0x%02X)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}
			}

		}

	}

	return break_flag
}
