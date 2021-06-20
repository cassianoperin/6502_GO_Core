package CONSOLE

import (
	"6502/CORE"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StartConsole() {

	var (
		current_PC  uint16 // Detect PC changes
		breakpoints []uint16
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nConsole:\n")
	fmt.Printf("---------------------\n")

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

			// for {
			// 	CORE.CPU_Interpreter()

			// 	if CORE.PC != current_PC {
			// 		printHeader()
			// 	}

			// 	fmt.Println()
			// }

			// Number os steps
			for {

				for CORE.PC == current_PC {
					CORE.CPU_Interpreter()
				}

				// Update the current_PC value once the opcode has changed
				current_PC = CORE.PC

				printHeader()
				fmt.Println()
			}

		} else { // Command not found
			fmt.Print("Command not found\n")
		}

	}

}

func printHelp() {
	fmt.Printf("\nquit\t\t\t\tQuit console\nhelp\t\t\t\tPrint help menu\nstep\t\t\t\tExecute current opcode\nadd_breakpoint <value>\t\tAdd a breakpoint\ndel_breakpoint <value>\t\tDelete a breakpoint\nshow_breakpoints\t\tShow breakpoints\nstep <value>\t\t\tExecute <value> opcodes\nrun\t\t\t\tRun the emulator\n\n")
}

func printHeader() {
	fmt.Printf("\tPC\tA\tX\tY\tSP\tNV-BDIZC\tCycle\n")
	fmt.Printf("\t%04X\t%02X\t%02X\t%02X\t%02X\t%d%d%d%d%d%d%d%d\t%d\n", CORE.PC, CORE.A, CORE.X, CORE.Y, CORE.SP, CORE.P[7], CORE.P[6], CORE.P[5], CORE.P[4], CORE.P[3], CORE.P[2], CORE.P[1], CORE.P[0], CORE.Cycle)
}

func RemoveIndex(s []uint16, index int) []uint16 {
	return append(s[:index], s[index+1:]...)
}
