package CONSOLE

import (
	"fmt"
	"strconv"
)

// Print Help Menu
func Console_Command_RunLimit(text_slice []string) {
	// Test the command syntax
	if len(text_slice) == 1 {

		// Show current value
		fmt.Printf("Current Run Limit = %d\n\n", run_limit)

	} else if len(text_slice) > 2 {

		// Print run_limit usage
		Console_PrintRunLimitErr()

	} else {

		// Convert string value to integer
		value, err := strconv.Atoi(text_slice[1])
		if err != nil {
			// handle error
			fmt.Printf("Invalid value: %s\n\n", text_slice[1])
		} else {
			run_limit = value
			fmt.Printf("New run limit = %d\n\n", run_limit)
		}

	}
}
