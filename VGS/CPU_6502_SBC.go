package VGS

import (
	"fmt"
	"strconv"
)

// SBC  Subtract Memory from Accumulator with Borrow (zeropage)
//
//      A - M - C -> A                   N Z C I D V
//                                       + + + - - +
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     SBC #oper     E9    2     2
//      zeropage      SBC oper      E5    2     3
//      zeropage,X    SBC oper,X    F5    2     4
//      absolute      SBC oper      ED    3     4
//      absolute,X    SBC oper,X    FD    3     4*
//      absolute,Y    SBC oper,Y    F9    3     4*
//      (indirect,X)  SBC (oper,X)  E1    2     6
//      (indirect),Y  SBC (oper),Y  F1    2     5*

func opc_SBC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Check for extra cycles (*) in the first opcode cycle
	if opc_cycle_count == 1 {
		if opcode == 0xFD || opcode == 0xF9 || opcode == 0xF1 {
			// Add 1 to cycles if page boundary is crossed
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

		// Original value of A and P0
		var (
			original_A        byte = A
			original_P0       byte = P[0]
			original_P7       byte = P[7]
			Mem_1s_complement byte = 255 - Memory[memAddr] // Memory value one's complement (bits inverted)
		)

		// --------------------------------- Binary / Hex Mode -------------------------------- //

		if P[3] == 0 {

			// Result
			// SBC is an ADC but with Memory value as one's complement (bits inverted)
			A = A + Mem_1s_complement + P[0]

			// Update the oVerflow flag
			flags_V(original_A, Mem_1s_complement, original_P0)

			// Update the carry flag value
			flags_C_ADC_SBC(original_A, Mem_1s_complement, original_P0)

			flags_Z(A)
			flags_N(A)

			// ----------------------------------- Decimal Mode ----------------------------------- //

		} else {

			var (
				bcd_Mem        int64
				tmp_A          int
				tmp_A_unsigned int
			)

			// Store the decimal value of the original A (hex)
			bcd_A, _ := strconv.ParseInt(fmt.Sprintf("%X", A), 0, 32)

			// Store the decimal value of the original Memory Address (hex)
			bcd_Mem, _ = strconv.ParseInt(fmt.Sprintf("%X", Memory[memAddr]), 0, 32)

			borrow := original_P0 ^ 1

			// Store the decimal result of A (must be trasformed in hex to be stored)
			tmp_A_unsigned = int(bcd_A) - int(bcd_Mem) - int(borrow)
			// BCD wrap-around between 0 and 99
			if tmp_A_unsigned < 0 {
				tmp_A = tmp_A_unsigned + 100
			} else {
				tmp_A = tmp_A_unsigned
			}

			// Convert the Decimal Result in to Hex to be returned to Accumulator
			bcd_Result, _ := strconv.ParseInt(fmt.Sprintf("%d", tmp_A), 16, 32)

			// Tranform the uint64 into a byte
			A = byte(bcd_Result)

			// ------------------------------ Flags ------------------------------ //

			// Update the oVerflow flag
			flags_V(original_A, Memory[memAddr], original_P0)

			// Update the carry flag value
			if tmp_A_unsigned >= 0x00 {
				P[0] = 1
			} else {
				P[0] = 0
			}
			if Debug {
				fmt.Printf("\tFlag C: %d -> %d\n", original_P7, P[7])
			}

			flags_Z(A)

			// Negative flag
			if tmp_A_unsigned < 0x00 {
				P[7] = 1
			} else {
				P[7] = 0
			}
			if Debug {
				fmt.Printf("\tFlag N: %d -> %d\n", original_P0, P[0])
			}

		}

		// --------------------------------------- Debug -------------------------------------- //

		if Debug {

			// Decimal flag OFF (Binary or Hex Mode)
			if P[3] == 0 {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow.\tA = A(%d) - Memory[%02X](%d) - Borrow(Inverted Carry)(%d) = %d\n", opcode, Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0^1, A)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow.\tA = A(%d) - Memory[%02X](%d) - Borrow(Inverted Carry)(%d) = %d\n", opcode, Memory[PC+2], Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0^1, A)
				}

				// Decimal flag ON (Decimal Mode)
			} else {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow. [Decimal Mode]\tA = A(%02X) - Memory[%02X](%02X) - Borrow(Inverted Carry)(%d) = %02X\n", opcode, Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0^1, A)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow. [Decimal Mode]\tA = A(%02X) - Memory[%02X](%02X) - Borrow(Inverted Carry)(%d) = %02X\n", opcode, Memory[PC+2], Memory[PC+1], mode, original_A, memAddr, Memory[memAddr], original_P0^1, A)
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
