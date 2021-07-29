package CONSOLE

import (
	"fmt"
	"strconv"
)

// Print Help Menu
func Console_Command_StepSkipDebug(text_slice []string) {
	// Test the command syntax
	if len(text_slice) == 1 {

		// Show current value
		fmt.Printf("Current Step Limit = %d\n\n", step_limit)

	} else if len(text_slice) > 2 {

		// Print step_limit usage
		Console_PrintStepLimitErr()

	} else {

		// Convert string value to integer
		value, err := strconv.Atoi(text_slice[1])
		if err != nil {
			// handle error
			fmt.Printf("Invalid value: %s\n\n", text_slice[1])
		} else {
			step_limit = value
			fmt.Printf("New step limit = %d\n\n", step_limit)
		}
	}
}
