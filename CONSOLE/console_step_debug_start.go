package CONSOLE

import (
	"fmt"
	"strconv"
)

// Console step_debug_start command
func Console_Command_StepDebugStart(text_slice []string) {
	// Test the command syntax
	if len(text_slice) == 1 {

		// Show current value
		fmt.Printf("Step debug messages starts on cycle = %d\n\n", step_debug_start)

	} else if len(text_slice) > 2 {

		// Print step_debug_start usage
		fmt.Printf("Usage: step_debug_start <Value>\n\n")

	} else {

		// Convert string value to integer
		value, err := strconv.Atoi(text_slice[1])
		if err != nil {
			// handle error
			fmt.Printf("Invalid value: %s\n\n", text_slice[1])
		} else {
			step_debug_start = uint64(value)
			fmt.Printf("Now Step command will start showing debug after cycle = %d\n\n", step_debug_start)
		}
	}
}
