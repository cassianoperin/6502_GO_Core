package CONSOLE

import (
	"fmt"
)

// Console show_breakpoint command
func Console_Command_ShowBreakpoint(text_slice []string) {

	// Test the command syntax
	if len(text_slice) == 1 {

		fmt.Printf("\n#\t\tLocation\tHEX\t\tDEC\n")

		for i := 0; i < len(breakpoints); i++ {
			fmt.Printf("Breakpoint %d:\t%s\t\t0x%02X\t\t%d\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
			if i == len(breakpoints)-1 {
				fmt.Println()
			}
		}

		if len(breakpoints) == 0 {
			fmt.Printf("No Breakpoint found.\n\n")
		}

	} else {

		// Print add_breakpoint usage
		fmt.Printf("show_breakpoints command doesn't accept arguments\n\n")

	}

}
