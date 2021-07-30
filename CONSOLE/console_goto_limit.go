package CONSOLE

import (
	"fmt"
	"strconv"
)

// Console goto_limit command
func Console_Command_GotoLimit(text_slice []string) {
	// Test the command syntax
	if len(text_slice) == 1 {

		// Show current value
		fmt.Printf("Current Goto Limit = %d\n\n", goto_limit)

	} else if len(text_slice) > 2 {

		// Print goto_limit usage
		Console_PrintGotoLimitErr()

	} else {

		// Convert string value to integer
		value, err := strconv.Atoi(text_slice[1])
		if err != nil {
			// handle error
			fmt.Printf("Invalid value: %s\n\n", text_slice[1])
		} else {
			goto_limit = value
			fmt.Printf("New goto limit = %d\n\n", goto_limit)
		}
	}
}
