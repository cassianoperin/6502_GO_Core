package CONSOLE

import (
	"fmt"
	"strconv"
)

// Console del_breakpoint command
func Console_Command_DelBreakpoint(text_slice []string) {

	// Test the command syntax
	if len(text_slice) == 1 || len(text_slice) > 2 {

		// Print add_breakpoint usage
		fmt.Printf("Usage: del_breakpoint <breakpoint number>\n")

	} else {

		value, err := strconv.Atoi(text_slice[1])
		if err != nil {
			// handle error
			fmt.Printf("Invalid value %s\n\n", text_slice[1])
		} else {
			if value < len(breakpoints) {
				breakpoints = Console_Remove_breakpoint(breakpoints, value)
				fmt.Printf("Breakpoint %d removed.\n\n", value)
			} else {
				fmt.Printf("Breakpoint not found\n\n")
			}

		}
	}

}
