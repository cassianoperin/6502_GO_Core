package CONSOLE

import (
	"6502/CORE"
	"fmt"
)

// Console disassemble command
func Console_Command_Disassemble(text_slice []string) {

	if len(text_slice) == 2 { // With ONE argument (show this memory value)

		fmt.Println()

		// Check if input is Decimar of Hexadecimal and convert to integer
		mem_arg, error_flag := Console_Hex_or_Dec(text_slice[1])

		if !error_flag {

			// Print Memory Value
			if mem_arg < 0 || mem_arg >= len(CORE.Memory) {
				fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFFFF (65536)\n\n")
			} else {
				// Disassemble Memory address
				print_debug_console(opcode_map, mem_arg)
			}

		} else {
			fmt.Printf("Invalid value %s\n\n", text_slice[1])
		}
		fmt.Println()

	} else if len(text_slice) == 3 { // With TWO argument (show memory value of the range)

		fmt.Println()

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

					var opc_bytes byte = 1

					// Get the number of bytes of the opcode to skip the operands
					for j := 0; j < len(opcode_map); j++ {

						if CORE.Memory[i] == opcode_map[j].code {
							opc_bytes = opcode_map[j].bytes
						}
					}

					// Disassemble Memory address
					print_debug_console(opcode_map, i)

					// Skip the operands based on the number of bytes of the last instruction
					i += int(opc_bytes - 1)
				}
				fmt.Println()
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
		fmt.Printf("Usage:\n   disassemble <address>\n   disassemble <start address> <end address>\n\n")
	}
}
