package CORE

import (
	"fmt"
)

// RTS  Return from Subroutine
//
//      pull PC, PC+1 -> PC              N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       RTS           60    1     6

func opc_RTS(bytes uint16, opc_cycles byte) {

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
			SP_Address = uint(SP)

			// 6502/6507 interpreter mode
		} else {
			// Stack is a 256-byte array whose location is hardcoded at page $01 ($0100-$01FF)
			SP_Address = uint(SP) + 256
		}

		PC = uint16(Memory[SP_Address+2])<<8 | uint16(Memory[SP_Address+1])

		// UNDOCUMENTED // Clear the addresses retrieved from the stack
		Memory[SP_Address+1] = 0
		Memory[SP_Address+2] = 0

		// Update the Stack Pointer (Increase the two values retrieved)
		SP += 2

		// Print Opcode Debug Message
		opc_RTS_DebugMsg(bytes)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_RTS_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tRTS  Return from Subroutine.\tPC = 0x%04X (+ 1 RTS instruction byte) = 0x%04X\n", opc_string, PC, PC+0x01)
		fmt.Println(dbg_show_message)
	}
}
