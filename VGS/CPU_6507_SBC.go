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

	// Inverted Carry
	var borrow byte = P[0] ^ 1

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

		original_A := A

		// --------------------------------- Binary / Hex Mode -------------------------------- //

		if P[3] == 0 {

			if Debug {
				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow.\tA = A(%d) - Memory[%02X](%d) - Borrow(Inverted Carry)(%d) = %d\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], borrow, A-Memory[memAddr]-borrow)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSBC  Subtract Memory from Accumulator with Borrow.\tA = A(%d) - Memory[%02X](%d) - Borrow(Inverted Carry)(%d) = %d\n", opcode, Memory[PC+2], Memory[PC+1], mode, A, memAddr, Memory[memAddr], borrow, A-Memory[memAddr]-borrow)
				}
				fmt.Println(dbg_show_message)
			}

			original_P0 := P[0]

			// Result
			// A = A - Memory[memAddr] - borrow
			A = A + (255 - Memory[memAddr]) + P[0]

			// For the flags:
			// The subtraction is VALUE1 (A) - VALUE2 (Memory[PC+1] - (P[0] ^ 1)
			// value2 := Memory[PC+1] - borrow

			// First V because it need the original carry flag value
			Flags_V_SBC(original_A, Memory[memAddr])
			// After, update the carry flag value
			// flags_C_Subtraction(original_A, A)

			// After, update the carry flag value
			// Set if overflow in bit 7 (the sum of values are smaller than original A)
			if A < original_A {
				P[0] = 1
				fmt.Println("Exit - ADC setou carry! Validar!")
				os.Exit(2)
			} else {
				P[0] = 0
			}

			if Debug {
				fmt.Printf("\tFlag C: %d -> %d\n", original_P0, P[0])
			}

			// // Clear Carry if overflow in bit 7 // NOT NECESSARY
			// if P[6] == 1 {
			// 	fmt.Printf("\n\tCarry cleared due to an overflow!")
			// 	P[0] = 0
			// }

			flags_Z(A)
			flags_N(A)

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1

			// Reset Opcode Extra Cycle counter
			opc_cycle_extra = 0

			// ----------------------------------- Decimal Mode ----------------------------------- //

		} else {
			fmt.Println("SBC DECIMAL NOT INPLEMENTED YET! EXITING")
			os.Exit(2)
		}
	}

}
