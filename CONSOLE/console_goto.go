package CONSOLE

import (
	"fmt"
)

// Print Help Menu
func Console_Command_Goto(text_slice []string) {

	// Test the command syntax
	if len(text_slice) == 1 || len(text_slice) > 2 {

		// Print goto usage
		fmt.Printf("Usage: goto <PC_ADDRESS>\n\n")

	} else {

		fmt.Println("goto com 1 arg")

	}

}
