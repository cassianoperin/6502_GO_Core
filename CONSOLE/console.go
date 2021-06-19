package CONSOLE

import (
	"6502/CORE"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartConsole() {

	var current_PC uint16 // Detect PC changes

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
			// CORE.CPU_Interpreter()
			// if CORE.PC != current_PC {
			// 	printHeader()
			// }

			for CORE.PC == current_PC {
				CORE.CPU_Interpreter()
			}
			printHeader()

		} else { // Command not found
			fmt.Print("Command not found\n")
		}

	}

}

func printHelp() {
	fmt.Printf("\nquit\tQuit console\nhelp\tPrint help menu\nXXX\txxxxxxx\n\n")
}

func printHeader() {
	fmt.Printf("Cycle: %d\tPC: 0x%04X\tA:%02X\tX:%02X\tY:%02X\tSP:%02X\n", CORE.Cycle, CORE.PC, CORE.A, CORE.X, CORE.Y, CORE.SP)
}
