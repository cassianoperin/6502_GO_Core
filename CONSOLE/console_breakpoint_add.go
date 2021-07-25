package CONSOLE

import (
	"fmt"
	"strings"
)

// Print Help Menu
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

				// // Test if the value start if 0x or 0X
				// if strings.HasPrefix(location_value[1], "0x") || strings.HasPrefix(location_value[1], "0X") {
				// 	// fmt.Println("seria HEX")

				// 	var hexaString string = location_value[1]
				// 	numberStr := strings.Replace(hexaString, "0x", "", -1)
				// 	numberStr = strings.Replace(numberStr, "0X", "", -1)

				// 	value, err := strconv.ParseInt(numberStr, 16, 64)

				// 	if err != nil {
				// 		fmt.Println("Invalid value.")
				// 	} else {
				// 		// Value limits
				// 		if location == "PC" {
				// 			if value <= 65535 && value >= 0 {
				// 				breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint16(value)})
				// 				fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
				// 			} else {
				// 				fmt.Println("Invalid value.")
				// 			}
				// 		}

				// 		if location == "A" || location == "X" || location == "Y" {
				// 			if value <= 255 && value >= 0 {
				// 				breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint16(value)})
				// 				fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
				// 			} else {
				// 				fmt.Println("Invalid value.")
				// 			}
				// 		}

				// 		if location == "CYCLE" {
				// 			if value >= 0 {
				// 				breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint16(value)})
				// 				fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
				// 			} else {
				// 				fmt.Println("Invalid value.")
				// 			}
				// 		}

				// 	}

				// } else {
				// 	// fmt.Println("seria DEC")

				// 	value, err := strconv.Atoi(location_value[1])

				// 	if err != nil {
				// 		fmt.Println("Invalid value.")
				// 	} else {
				// 		// Value limits
				// 		if location == "PC" {
				// 			if value <= 65535 && value >= 0 {
				// 				breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint16(value)})
				// 				fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
				// 			} else {
				// 				fmt.Println("Invalid value.")
				// 			}
				// 		}

				// 		if location == "A" || location == "X" || location == "Y" {
				// 			if value <= 255 && value >= 0 {
				// 				breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint16(value)})
				// 				fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
				// 			} else {
				// 				fmt.Println("Invalid value.")
				// 			}
				// 		}

				// 		if location == "CYCLE" {
				// 			if value >= 0 {
				// 				breakpoints = append(breakpoints, breakpoint{strings.ToUpper(location_value[0]), uint16(value)})
				// 				fmt.Printf("Breakpoint %d created.\n\n", len(breakpoints)-1)
				// 			} else {
				// 				fmt.Println("Invalid value.")
				// 			}
				// 		}

				// 	}
				// }
			} else {

				// Print add_breakpoint usage
				Console_PrintAddBrkErr()
			}

		}

	}

}
