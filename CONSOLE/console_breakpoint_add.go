package CONSOLE

import (
	"fmt"
	"strings"
)

// Console add_breakpoint command
func Console_Command_AddBreakpoint(text_slice []string) {

	var location_value []string

	// Test the command syntax
	if len(text_slice) == 1 || len(text_slice) > 2 {

		// Print add_breakpoint usage
		Console_PrintAddBrkErr()

	} else {

		// After, split the argument in LOCATION=VALUE
		location_value = strings.Split(text_slice[1], "=")

		if len(location_value) == 1 || len(location_value) > 2 || location_value[1] == "" || location_value[0] == "" {

			// Print add_breakpoint usage
			Console_PrintAddBrkErr()

		} else {

			location := strings.ToUpper(location_value[0])

			// Validate the value of locations
			if location == "PC" || location == "A" || location == "X" || location == "Y" || location == "CYCLE" {

				// Check if input is Decimar of Hexadecimal and convert to integer
				mem_arg, error_flag := Console_Hex_or_Dec(location_value[1])

				if !error_flag {
					// Value limits
					if location == "PC" {
						if mem_arg <= 65535 && mem_arg >= 0 {
							breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint64(mem_arg)})
							fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFFFF (65536)\n\n")
						}
					}

					if location == "A" || location == "X" || location == "Y" {
						if mem_arg <= 255 && mem_arg >= 0 {
							breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint64(mem_arg)})
							fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFF (255)\n\n")
						}
					}

					if location == "CYCLE" {
						if mem_arg >= 0 {
							breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint64(mem_arg)})
							fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
						} else {
							fmt.Println("Invalid value.")
						}
					}
				} else {
					fmt.Printf("Invalid value %s\n\n", location_value[1])
				}

			} else {

				// Print add_breakpoint usage
				Console_PrintAddBrkErr()
			}

		}

	}

}
