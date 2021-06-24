package CONSOLE

import (
	"6502/CORE"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instructuction struct {
	code        byte
	bytes       byte
	cycles      byte
	description string
	memory_mode string
}

func StartConsole() {

	var (
		current_PC  uint16 // Detect PC changes
		breakpoints []uint16
		opcode_map  []instructuction
	)

	opcode_map = []instructuction{
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

	// fmt.Println(opcode_map)
	// fmt.Println(opcode_map[1].cycles)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nConsole:\n\n")

	// Reset system
	CORE.Reset()

	// Print Header
	printHeader()

	for {

		current_PC = CORE.PC

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		// Exit
		if strings.Compare("quit", text) == 0 || strings.Compare("Quit", text) == 0 || strings.Compare("exit", text) == 0 || strings.Compare("Exit", text) == 0 {
			fmt.Printf("Exiting console.\n")
			os.Exit(0)

		} else if strings.Compare("help", text) == 0 || strings.Compare("Help", text) == 0 { // Help
			printHelp()

		} else if strings.Compare("", text) == 0 { // ENTER
			fmt.Print("-> \n")

		} else if strings.Compare("step", text) == 0 { // STEP

			// Print the opcode debug
			print_debug_console(opcode_map)

			for CORE.PC == current_PC {
				CORE.CPU_Interpreter()
			}
			printHeader()

		} else if strings.HasPrefix(text, "step ") { // STEP WITH ARGS

			// Convert string value to integer
			value, err := strconv.Atoi(strings.TrimPrefix(text, "step "))
			if err != nil {
				// handle error
				fmt.Printf("Invalid value: %s\n", strings.TrimPrefix(text, "step "))
			} else {

				// Number os steps
				for i := 0; i < value; i++ {

					// Print the opcode debug
					print_debug_console(opcode_map)

					for CORE.PC == current_PC {
						CORE.CPU_Interpreter()
					}

					// Update the current_PC value once the opcode has changed
					current_PC = CORE.PC

					printHeader()
					fmt.Println()
				}

			}
		} else if strings.HasPrefix(text, "add_breakpoint") { // ADD BREAKPOINT

			if len(strings.Split(text, " ")) == 1 || len(strings.Split(text, " ")) > 2 {
				fmt.Println("\nUsage: add_breakpoint <address>")
			} else {

				tmp_string := strings.Split(text, " ")

				value, err := strconv.Atoi(tmp_string[1])
				if err != nil {
					// handle error
					fmt.Printf("Invalid value %s\n", tmp_string[1])
				} else {
					if value <= 65535 && value >= 0 {
						breakpoints = append(breakpoints, uint16(value))
						fmt.Printf("Breakpoint %d created.\n", len(breakpoints)-1)
					} else {
						fmt.Println("Invalid address.")
					}

				}
			}

		} else if strings.HasPrefix(text, "del_breakpoint") { // DELETE BREAKPOINT

			if len(strings.Split(text, " ")) == 1 || len(strings.Split(text, " ")) > 2 {
				fmt.Println("\nUsage: del_breakpoint <breakpoint number>")
			} else {

				tmp_string := strings.Split(text, " ")

				value, err := strconv.Atoi(tmp_string[1])
				if err != nil {
					// handle error
					fmt.Printf("Invalid value %s\n", tmp_string[1])
				} else {
					if value < len(breakpoints) {
						breakpoints = RemoveIndex(breakpoints, value)
						fmt.Printf("Breakpoint %d removed.\n", len(breakpoints)-1)
					} else {
						fmt.Printf("Breakpoint not found\n")
					}

				}
			}

		} else if strings.Compare("show_breakpoints", text) == 0 { // SHOW BREAKPOINTS

			for i := 0; i < len(breakpoints); i++ {
				fmt.Printf("Breakpoint %d: %d\n", i, breakpoints[i])
			}

			if len(breakpoints) == 0 {
				fmt.Printf("No Breakpoint found.\n")
			}

		} else if strings.Compare("run", text) == 0 { // RUN

			for { // Add breakpoint control here

				select {
				case <-CORE.Second_timer: // Show the header and debug each second
					// Print the opcode debug
					print_debug_console(opcode_map)

					// Execute a new cycle
					for CORE.PC == current_PC {
						CORE.CPU_Interpreter()
					}

					// Onde finished the opcode (opcode changed), update the new current_PC
					current_PC = CORE.PC

					printHeader()
					fmt.Println()

				default: // Just run the CPU

					if CORE.Cycle < 96000000 {

						// Execute a new cycle
						for CORE.PC == current_PC {
							CORE.CPU_Interpreter()
						}

						// Onde finished the opcode (opcode changed), update the new current_PC
						current_PC = CORE.PC
					} else {
						// Print the opcode debug
						print_debug_console(opcode_map)

						// Execute a new cycle
						for CORE.PC == current_PC {
							CORE.CPU_Interpreter()
						}

						// Onde finished the opcode (opcode changed), update the new current_PC
						current_PC = CORE.PC

						printHeader()
						fmt.Println()
					}

				}

				// // Print the opcode debug
				// print_debug_console(opcode_map)

				// // Execute a new cycle
				// for CORE.PC == current_PC {
				// 	CORE.CPU_Interpreter()
				// }

				// // Onde finished the opcode (opcode changed), update the new current_PC
				// current_PC = CORE.PC

				// printHeader()
				// fmt.Println()
			}

		} else { // Command not found
			fmt.Print("Command not found\n")
		}

	}

}

func print_debug_console(opcode_map []instructuction) {

	// Search for the current opcode inside opcode map
	for i := 0; i < len(opcode_map); i++ {

		if CORE.Memory[CORE.PC] == opcode_map[i].code {
			opc_string, opc_operand := CORE.Debug_decode_console(opcode_map[i].bytes)
			fmt.Printf("\n\t$%04X\t%s %s\t%s %s (%s)\n\n", CORE.PC, opc_string, opc_operand, opcode_map[i].description, opc_operand, opcode_map[i].memory_mode)
		}

	}
}

func printHelp() {
	fmt.Printf("\nquit\t\t\t\tQuit console\nhelp\t\t\t\tPrint help menu\nstep\t\t\t\tExecute current opcode\nadd_breakpoint <value>\t\tAdd a breakpoint\ndel_breakpoint <value>\t\tDelete a breakpoint\nshow_breakpoints\t\tShow breakpoints\nstep <value>\t\t\tExecute <value> opcodes\nrun\t\t\t\tRun the emulator\n\n")
}

func printHeader() {
	fmt.Printf("--------------------------------------------------------------------------\n")
	fmt.Printf("\tPC\tA\tX\tY\tSP\tNV-BDIZC\tCycle\n")
	fmt.Printf("\t%04X\t%02X\t%02X\t%02X\t%02X\t%d%d%d%d%d%d%d%d\t%d\n", CORE.PC, CORE.A, CORE.X, CORE.Y, CORE.SP, CORE.P[7], CORE.P[6], CORE.P[5], CORE.P[4], CORE.P[3], CORE.P[2], CORE.P[1], CORE.P[0], CORE.Cycle)
}

func RemoveIndex(s []uint16, index int) []uint16 {
	return append(s[:index], s[index+1:]...)
}
