package CONSOLE

import (
	"6502/CORE"
	"fmt"
)

// Console debug command
func Console_Command_Debug(text_slice []string) {
	// Test the command syntax
	if len(text_slice) == 1 {

		// Show current value
		if CORE.Debug {
			fmt.Printf("Debug status: Enabled\n\n")
		} else {
			fmt.Printf("Debug status: Disabled\n\n")
		}

	} else if len(text_slice) > 2 {

		// Print debug usage
		fmt.Printf("Usage: debug <on|off>\n\n")

	} else {

		if text_slice[1] == "on" || text_slice[1] == "off" {
			if text_slice[1] == "on" {
				CORE.Debug = true
				fmt.Printf("Debug mode enabled\n\n")
			} else {
				CORE.Debug = false
				fmt.Printf("Debug mode disabled\n\n")
			}
		} else {
			// Print debug usage
			fmt.Printf("Usage: debug <on|off>\n\n")
		}

	}
}
