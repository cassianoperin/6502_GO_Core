package CORE

import "fmt"

// RLA  ROL + AND  (Unofficial)
//
//      C <- [76543210] <- C             N Z C I D V
//      A AND M -> A                     + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      RLA oper      27    2     5
//
// Operation: Rotate one bit left in memory, then AND accumulator with memory.
//
// Example:
// 			RLA $FC,X ;37 FC
// Equivalent Instructions:
// 			ROL $FC,X
// 			AND $FC,X

func opc_U_RLA(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s] [Unnoficial!!!]\tRLA  ROL + AND.\n", opcode, Memory[PC+1], mode)
			fmt.Println(dbg_show_message)
		}

		// ----------------------------- ROL ----------------------------- //

		// Original Carry Value
		carry_orig := P[0]

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tROL  Rotate One Bit Left.\tMemory[%d] Roll Left 1 bit\t(%d)\n", memAddr, (Memory[memAddr]<<1)+carry_orig)
			fmt.Println(dbg_show_message)
		}

		// Calculate the original bit7 and save it as the new Carry
		P[0] = Memory[memAddr] & 0x80 >> 7

		// Shift left the byte and put the original bit7 value in bit 1 to make the complete ROL
		Memory[memAddr] = (Memory[memAddr] << 1) + carry_orig

		flags_N(Memory[memAddr])
		flags_Z(Memory[memAddr])
		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", carry_orig, P[0])
		}

		// ----------------------------- AND ----------------------------- //

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\n\tAND  AND Memory with Accumulator.\tA = A(%d) & Memory[%02X](%d)\t(%d)\n", A, memAddr, Memory[memAddr], A&Memory[memAddr])
			fmt.Println(dbg_show_message)
		}

		A = A & Memory[memAddr]

		flags_Z(A)
		flags_N(A)

		// ----------------------------- Common ----------------------------- //

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
