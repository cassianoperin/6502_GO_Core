package CONSOLEMODE

import (
	"6502/CORE"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartConsole() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Console:")
	fmt.Println("---------------------")

	// Reset system
	CORE.Reset()

	for {

		fmt.Printf("PC: %04X\n", CORE.PC)
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
		} else { // Command not found
			printHelp()
		}

	}

}

func printHelp() {
	fmt.Printf("exit\tQuit cosole.\t\tquit\tQuit console\t\thelp\tPrint help menu\nXXX\txxxxxxx")
}
