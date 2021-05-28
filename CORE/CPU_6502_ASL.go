package CORE

import (
	"fmt"
)

// ASL  Shift Left One Bit (Memory or Accumulator)
//
//      C <- [76543210] <- 0             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   ASL A         0A    1     2
//      zeropage      ASL oper      06    2     5
//      zeropage,X    ASL oper,X    16    2     6
//      absolute      ASL oper      0E    3     6
//      absolute,X    ASL oper,X    1E    3     7

// ASL shifts all bits left one position. 0 is shifted into bit 0 and the original bit 7 is shifted into the Carry.

// ------------------------------------ Accumulator ------------------------------------ //

func opc_ASL_A(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Print Opcode Debug Message
		opc_ASL_A_DebugMsg(bytes)

		P[0] = A >> 7

		A = A << 1

		flags_N(A)
		flags_Z(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

func opc_ASL_A_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Accumulator]\tASL  Shift Left One Bit.\tA = A(%d) Shift Left 1 bit\t(%d).\tCarry (Original A bit 7): %d\n", opc_string, A, A<<1, A>>7)
		fmt.Println(dbg_show_message)
	}
}

// --------------------------------------- Memory -------------------------------------- //

func opc_ASL(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Print Opcode Debug Message
		opc_ASL_DebugMsg(bytes, mode, memAddr)

		P[0] = Memory[memAddr] >> 7

		Memory[memAddr] = Memory[memAddr] << 1

		flags_N(Memory[memAddr])
		flags_Z(Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}
}

func opc_ASL_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tASL  Shift Left One Bit.\tMemory[0x%02X]: (%d) Shift Left 1 bit\t(%d).\tCarry (Original Memory address bit 7): %d\n", opc_string, mode, memAddr, Memory[memAddr], Memory[memAddr]<<1, Memory[memAddr]>>7)
		fmt.Println(dbg_show_message)
	}
}
