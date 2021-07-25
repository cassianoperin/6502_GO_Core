package CONSOLE

import (
	"6502/CORE"
	"fmt"
)

// Print Help Menu
func Console_Command_Mem(text_slice []string) {

	if len(text_slice) == 1 { // Without arguments (show all memory)
		fmt.Printf("\t00 01 02 03 04 05 06 07 08 09 0A 0B 0C 0D 0E 0F\n")
		fmt.Printf("\t-----------------------------------------------")
		for i := 0; i < len(CORE.Memory); i++ {

			// Break lines
			if i%16 == 0 {
				fmt.Printf("\n%04X\t", i)
			}

			// Print memory
			fmt.Printf("%02X ", CORE.Memory[i])

		}
		fmt.Println()

	} else if len(text_slice) == 2 { // With ONE argument (show this memory value)

		// Check if input is Decimar of Hexadecimal and convert to integer
		mem_arg, error_flag := Console_Hex_or_Dec(text_slice[1])

		if !error_flag {

			// Print Memory Value
			if mem_arg < 0 || mem_arg >= len(CORE.Memory) {
				fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFFFF (65536)\n\n")
			} else {
				fmt.Printf("%02X\n\n", CORE.Memory[mem_arg])
			}

		} else {
			fmt.Printf("Invalid value %s\n\n", text_slice[1])
		}

	} else if len(text_slice) == 3 { // With TWO argument (show memory value of the range)

		// Check if input is Decimar of Hexadecimal and convert to integer
		mem_arg, error_flag := Console_Hex_or_Dec(text_slice[1])
		mem_arg2, error_flag2 := Console_Hex_or_Dec(text_slice[2])

		if !error_flag && !error_flag2 {
			// Invald Address Range
			if mem_arg < 0 || mem_arg >= len(CORE.Memory) {
				fmt.Printf("Invalid Address in first argument. Should be in range 0x0000 and 0xFFFF (65536)\n\n")
				error_flag = true
			} else if mem_arg2 < 0 || mem_arg2 >= len(CORE.Memory) {
				fmt.Printf("Invalid Address in second argument. Should be in range 0x0000 and 0xFFFF (65536)\n\n")
				error_flag = true
			} else if mem_arg > mem_arg2 {
				fmt.Printf("Start address should be less or equal end address\n\n")
				error_flag = true
			} else { // Print Memory Value
				for i := mem_arg; i <= mem_arg2; i++ {
					fmt.Printf("%02X ", CORE.Memory[i])
				}
				fmt.Printf("\n\n")
			}
		} else {
			// Invalid values
			if error_flag && error_flag2 {
				fmt.Printf("Invalid values %s and %s\n\n", text_slice[1], text_slice[2])
			} else if error_flag {
				fmt.Printf("Invalid value %s\n\n", text_slice[1])
			} else if error_flag2 {
				fmt.Printf("Invalid value %s\n\n", text_slice[2])
			}
		}

	} else {
		fmt.Printf("Usage:\n   mem\n   mem <address>\n   mem <start address> <end address>\n\n")
	}
}
