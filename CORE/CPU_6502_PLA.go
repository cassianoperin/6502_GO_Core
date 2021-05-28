package CORE

import (
	"fmt"
)

// PLA  Pull Accumulator from Stack
//
//      pull A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PLA           68    1     4

func opc_PLA(bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		var SP_Address uint

		// Atari 2600 interpreter mode
		if CPU_MODE == 0 {
			SP_Address = uint(SP + 1)

			// 6502/6507 interpreter mode
		} else {
			// Stack is a 256-byte array whose location is hardcoded at page $01 ($0100-$01FF)
			SP_Address = uint(SP+1) + 256
		}

		A = Memory[SP_Address]

		// Print Opcode Debug Message
		opc_PLA_DebugMsg(bytes, SP_Address)

		flags_N(A)
		flags_Z(A)

		SP++

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_PLA_DebugMsg(bytes uint16, SP_Address uint) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tPLA  Pull Accumulator from Stack.\tA = Memory[0x%02X] (%d) | SP++\n", opc_string, SP_Address, A)
		fmt.Println(dbg_show_message)
	}
}
