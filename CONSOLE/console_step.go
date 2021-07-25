package CONSOLE

import (
	"fmt"
	"strconv"
)

// Print Help Menu
func Console_Command_Step(text_slice []string) {

	if len(text_slice) == 1 {
		// Execute one instruction
		Console_Step(opcode_map)

	} else if len(text_slice) == 2 {

		var breakpoint_flag bool

		// Convert string value to integer
		value, err := strconv.Atoi(text_slice[1])
		if err != nil {
			// handle error
			fmt.Printf("Invalid value: %s\n\n", text_slice[1])
		} else {

			// Number os steps
			if value <= step_limit {
				for i := 0; i < value; i++ {

					// Execute one instruction
					Console_Step(opcode_map)

					// Check Breakpoints
					breakpoint_flag = Console_Check_breakpoints(breakpoint_flag)

					// Exit for loop if breakpoint has been found
					if breakpoint_flag {
						break
					}

				}
			} else {
				fmt.Printf("Current step limit = %d\n\n", step_limit)
			}
		}

	} else {
		fmt.Printf("Usage: step <number of cycles>\n\n")
	}

}
