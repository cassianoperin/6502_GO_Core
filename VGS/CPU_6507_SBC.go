package VGS

import (
	"fmt"
	"os"
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
			Mem_1s_complement byte = 255 - Memory[memAddr] // Memory value one's complement (bits inverted)
		)

		// --------------------------------- Binary / Hex Mode -------------------------------- //

		if P[3] == 0 {

			if Debug {

			}

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
			fmt.Println("SBC DECIMAL NOT INPLEMENTED YET! EXITING")
			os.Exit(2)
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

				fmt.Printf("Implement Decimal mode sbc debug messages\n")
				os.Exit(2)

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
