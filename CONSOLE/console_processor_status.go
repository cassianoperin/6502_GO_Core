package CONSOLE

import (
	"6502/CORE"
	"fmt"
	"strings"
)

// Console processor_status command
func Console_Command_ProcessorStatus(text_slice []string) {

	var location_value []string

	// Test the command syntax
	if len(text_slice) == 1 || len(text_slice) > 2 {

		// Print usage
		fmt.Printf("Usage: processor_status <N|V|B|D|I|Z|C> = <0|1>\n\n")
	} else {

		// After, split the argument in LOCATION=VALUE
		location_value = strings.Split(text_slice[1], "=")

		if len(location_value) == 1 || len(location_value) > 2 || location_value[1] == "" || location_value[0] == "" {

			// Print usage
			fmt.Printf("Usage: processor_status <N|V|B|D|I|Z|C> = <0|1>\n\n")

		} else {

			location := strings.ToUpper(location_value[0])

			// Validate the value of locations
			if location == "N" || location == "V" || location == "B" || location == "D" || location == "I" || location == "Z" || location == "C" {

				// Check if input is Decimar of Hexadecimal and convert to integer
				mem_arg, error_flag := Console_Hex_or_Dec(location_value[1])

				if !error_flag {
					// Value limits
					if location == "N" {
						if mem_arg <= 1 && mem_arg >= 0 {
							CORE.P[7] = byte(mem_arg)
							fmt.Printf("\tP[7] - Negative flag set to %d\n\n", CORE.P[7])
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid value. Should be 0 or 1.\n\n")
						}
					} else if location == "V" {
						if mem_arg <= 1 && mem_arg >= 0 {
							CORE.P[6] = byte(mem_arg)
							fmt.Printf("\tP[6] - Overflow flag set to %d\n\n", CORE.P[6])
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid value. Should be 0 or 1.\n\n")
						}
					} else if location == "B" {
						if mem_arg <= 1 && mem_arg >= 0 {
							CORE.P[4] = byte(mem_arg)
							fmt.Printf("\tP[4] - Break flag set to %d\n\n", CORE.P[4])
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid value. Should be 0 or 1.\n\n")
						}
					} else if location == "D" {
						if mem_arg <= 1 && mem_arg >= 0 {
							CORE.P[3] = byte(mem_arg)
							fmt.Printf("\tP[3] - Decimal flag set to %d\n\n", CORE.P[3])
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid value. Should be 0 or 1.\n\n")
						}
					} else if location == "I" {
						if mem_arg <= 1 && mem_arg >= 0 {
							CORE.P[2] = byte(mem_arg)
							fmt.Printf("\tP[2] - IRQ Disable flag set to %d\n\n", CORE.P[2])
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid value. Should be 0 or 1.\n\n")
						}
					} else if location == "Z" {
						if mem_arg <= 1 && mem_arg >= 0 {
							CORE.P[1] = byte(mem_arg)
							fmt.Printf("\tP[1] - Zero flag set to %d\n\n", CORE.P[1])
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid value. Should be 0 or 1.\n\n")
						}
					} else if location == "C" {
						if mem_arg <= 1 && mem_arg >= 0 {
							CORE.P[0] = byte(mem_arg)
							fmt.Printf("\tP[0] - Carry flag set to %d\n\n", CORE.P[0])
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid value. Should be 0 or 1.\n\n")
						}
					}

				} else {
					fmt.Printf("Invalid value %s\n\n", location_value[1])
				}

			} else {

				// Print usage
				fmt.Printf("Usage: processor_status <N|V|B|D|I|Z|C> = <0|1>\n\n")
			}

		}

	}

}
