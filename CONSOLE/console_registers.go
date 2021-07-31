package CONSOLE

import (
	"6502/CORE"
	"fmt"
	"strings"
)

// Console registers command
func Console_Command_Registers(text_slice []string) {

	var location_value []string

	// Test the command syntax
	if len(text_slice) == 1 || len(text_slice) > 2 {

		// Print usage
		fmt.Printf("Usage: registers <A|X|Y|PC>=<value>\n\n")
	} else {

		// After, split the argument in LOCATION=VALUE
		location_value = strings.Split(text_slice[1], "=")

		if len(location_value) == 1 || len(location_value) > 2 || location_value[1] == "" || location_value[0] == "" {

			// Print usage
			fmt.Printf("Usage: registers <A|X|Y|PC>=<value>\n\n")

		} else {

			location := strings.ToUpper(location_value[0])

			// Validate the value of locations
			if location == "PC" || location == "A" || location == "X" || location == "Y" {

				// Check if input is Decimar of Hexadecimal and convert to integer
				mem_arg, error_flag := Console_Hex_or_Dec(location_value[1])

				if !error_flag {
					// Value limits
					if location == "PC" {
						if mem_arg <= 65535 && mem_arg >= 0 {
							CORE.PC = uint16(mem_arg)
							fmt.Printf("\tPC set to 0x%02X (%d)\n\n", CORE.PC, CORE.PC)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFFFF (65536)\n\n")
						}
					}

					if location == "A" {
						if mem_arg <= 255 && mem_arg >= 0 {
							CORE.A = byte(mem_arg)
							fmt.Printf("\tA set to 0x%02X (%d)\n\n", CORE.A, CORE.A)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFF (255)\n\n")
						}
					}

					if location == "X" {
						if mem_arg <= 255 && mem_arg >= 0 {
							CORE.X = byte(mem_arg)
							fmt.Printf("\tX set to 0x%02X (%d)\n\n", CORE.X, CORE.X)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFF (255)\n\n")
						}
					}

					if location == "Y" {
						if mem_arg <= 255 && mem_arg >= 0 {
							CORE.Y = byte(mem_arg)
							fmt.Printf("\tY set to %02X (%d)\n\n", CORE.Y, CORE.Y)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFF (255)\n\n")
						}
					}

				} else {
					fmt.Printf("Invalid value %s\n\n", location_value[1])
				}

			} else {

				// Print usage
				fmt.Printf("Usage: registers <A|X|Y|PC>=<value>\n\n")
			}

		}

	}

}
