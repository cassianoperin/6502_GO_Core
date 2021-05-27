package CORE

import (
	"fmt"
	"strconv"
)

// ADC  Add Memory to Accumulator with Carry (zeropage)
//
//      A + M + C -> A, C                N Z C I D V
//     	                                 + + + - - +
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate	  ADC #oper	    69    2     2
//      zeropage      ADC oper      65    2     3
//      zeropage,X    ADC oper,X    75    2     4
//      absolute      ADC oper      6D    3     4
//      absolute,X    ADC oper,X    7D    3     4*
//      absolute,Y    ADC oper,Y    79    3     4*
//      (indirect,X)  ADC (oper,X)  61    2     6
//      (indirect),Y  ADC (oper),Y  71    2     5*

func opc_ADC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles+opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles+opc_cycle_extra {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Original value of A and P0
		var (
			original_A  byte = A
			original_P0 byte = P[0]
		)

		// --------------------------------- Binary / Hex Mode -------------------------------- //

		if P[3] == 0 {

			A = A + Memory[memAddr] + P[0]

			// Update the oVerflow flag
			flags_V(original_A, Memory[memAddr], original_P0)

			// Update the carry flag value
			flags_C_ADC_SBC(original_A, Memory[memAddr], original_P0)

			flags_Z(A)
			flags_N(A)

			// ----------------------------------- Decimal Mode ----------------------------------- //

		} else {

			var bcd_Mem int64

			// Store the decimal value of the original A (hex)
			bcd_A, _ := strconv.ParseInt(fmt.Sprintf("%X", A), 0, 32)

			// Store the decimal value of the original Memory Address (hex)
			bcd_Mem, _ = strconv.ParseInt(fmt.Sprintf("%X", Memory[memAddr]), 0, 32)

			// Store the decimal result of A (must be trasformed in hex to be stored)
			tmp_A := byte(bcd_A) + byte(bcd_Mem) + P[0]

			// Convert the Decimal Result in to Hex to be returned to Accumulator
			bcd_Result, _ := strconv.ParseInt(fmt.Sprintf("%d", tmp_A), 16, 32)

			// Tranform the uint64 into a byte (if > 255 will be rotated)
			A = byte(bcd_Result)

			// ------------------------------ Flags ------------------------------ //

			// Update the oVerflow flag
			flags_V(original_A, Memory[memAddr], original_P0)

			// Update the carry flag value
			if bcd_Result > 0x99 {
				P[0] = 1
			} else {
				P[0] = 0
			}
			// fmt.Printf("P[0] novo (sem contar negativo): %d\n\n", P[0])
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
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Binary/Hex Mode]\tA = A(%d) + Memory[%02X](%d) + Carry (%d)) = %d\n", opcode, Memory[PC+2], Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0, A)
				}

				// Decimal flag ON (Decimal Mode)
			} else {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Decimal Mode]\tA = A(%02x) + Memory[%02X](%02x) + Carry (%02x)) = %02X\n", opcode, Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0, A)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Decimal Mode]\tA = A(%02x) + Memory[%02X](%02x) + Carry (%d)) = %02X\n", opcode, Memory[PC+2], Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0, A)
				}

			}

			fmt.Println(dbg_show_message)

		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0

	}

}
