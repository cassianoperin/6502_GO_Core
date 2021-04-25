package VGS

import (
	"fmt"
	"strconv"
)

// ADC  Add Memory to Accumulator with Carry (zeropage)
//
// 	A + M + C -> A, C                N Z C I D V
// 	                          	   + + + - - +
//
// 	addressing    assembler    opc  bytes  cyles
// 	--------------------------------------------
// 	zeropage      ADC oper      65    2     3
//	absolute,X    ADC oper,X    7D    3     4*
//	immediate	  ADC #oper	    69    2     2

func opc_ADC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Check for extra cycles (*) in the first opcode cycle
	if opc_cycle_count == 1 {
		if opcode == 0x7D {
			// Add 1 to cycles if page boundery is crossed
			if MemPageBoundary(memAddr, PC) {
				opc_cycle_extra = 1
			}
		}
	}

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles+opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles+opc_cycle_extra {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Original value of A and P0 to be used on debug messages
		original_A := A
		original_P0 := P[0]

		// --------------------------------- Binary / Hex Mode -------------------------------- //

		if P[3] == 0 {

			A = A + Memory[memAddr] + P[0]

			// ------------------------------ Flags ------------------------------ //
			// First V because it need the original carry flag value
			Flags_V_ADC(original_A, Memory[memAddr])

			// After, update the carry flag value
			flags_C(original_A, A)
			flags_Z(A)
			flags_N(A)

			// ----------------------------------- Decimal Mode ----------------------------------- //

		} else {
			var bcd_Mem int64

			// Store the decimal value of the original A (hex)
			bcd_A, _ := strconv.ParseInt(fmt.Sprintf("%02X", A), 0, 32)

			// Immediate memory mode
			// Store the decimal value of the original Memory Address (hex)
			bcd_Mem, _ = strconv.ParseInt(fmt.Sprintf("%02X", Memory[memAddr]), 0, 32)

			// Store the decimal result of A (must be trasformed in hex to be stored)
			tmp_A := byte(bcd_A) + byte(bcd_Mem) + P[0]

			// Convert the Decimal Result in to Hex to be returned to Accumulator
			bcd_Result, _ := strconv.ParseInt(fmt.Sprintf("%d", tmp_A), 16, 32)

			// Tranform the uint64 into a byte (if > 255 will be rotated)
			A = byte(bcd_Result)

			// ------------------------------ Flags ------------------------------ //

			// Immediate memory mode
			// First V because it need the original carry flag value
			Flags_V_ADC(original_A, Memory[memAddr])

			// After, update the carry flag value
			// For Decimal Mode works different, if the sum of the values is > 255, set it
			if bcd_Result > 255 {
				P[0] = 1
			} else {
				P[0] = 0
			}
			// Show Carry debug
			if Debug {
				fmt.Printf("\tFlag C: %d -> %d\n", original_P0, P[0])
			}
			flags_Z(A)
			flags_N(A)

		}

		// --------------------------------------- Debug -------------------------------------- //

		if Debug {

			// Decimal flag OFF (Binary or Hex Mode)
			if P[3] == 0 {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Binary/Hex Mode]\tA = A(%d) + Memory[%02X](%d) + Carry (%d)) = %d\n", opcode, Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0, A)
					fmt.Println(dbg_show_message)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Binary/Hex Mode]\tA = A(%d) + Memory[%02X](%d) + Carry (%d)) = %d\n", opcode, Memory[PC+2], Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0, A)
					fmt.Println(dbg_show_message)
				}

				// Decimal flag ON (Decimal Mode)
			} else {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Decimal Mode]\tA = A(%02x) + Memory[%02X](%02x) + Carry (%02x)) = %02X\n", opcode, Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0, A)
					fmt.Println(dbg_show_message)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Decimal Mode]\tA = A(%02x) + Memory[%02X](%02x) + Carry (%d)) = %02X\n", opcode, Memory[PC+2], Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0, A)
					fmt.Println(dbg_show_message)
				}

			}

		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0

	}

}
