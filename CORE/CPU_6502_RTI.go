package CORE

import (
	"fmt"
)

// RTI  Return from Interrupt
//
//      pull SR, pull PC
//
//                                      N Z C I D V
//                                      from stack
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       RTI           40    1    6

// Order
// Restore P
// Restore PC(lo)
// Restore PC(hi)

func opc_RTI(bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// ---------- Restore P ---------- //

		var SP_Address uint

		// Atari 2600 interpreter mode
		if CPU_MODE == 0 {
			SP_Address = uint(SP + 1)

			// 6502/6507 interpreter mode
		} else {
			// Stack is a 256-byte array whose location is hardcoded at page $01 ($0100-$01FF)
			SP_Address = uint(SP+1) + 256
		}

		// Turn the stack value into the processor status
		for i := 0; i < len(P); i++ {

			// The B Flag, PLP and RTI pull a byte from the stack and set all the flags. They ignore bits 5 and 4.
			if i == 4 || i == 5 {
				// P[i] = 1
				// Just ignore both
			} else {
				P[i] = (Memory[SP_Address] >> i) & 0x01
			}

		}

		SP++

		// ---------- Restore PC ---------- //

		// Read the Opcode from PC+1 and PC bytes (Little Endian)
		PC = uint16(Memory[SP_Address+2])<<8 | uint16(Memory[SP_Address+1])

		SP += 2

		// Print Opcode Debug Message
		opc_RTI_DebugMsg(bytes, SP_Address)

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_RTI_DebugMsg(bytes uint16, SP_Address uint) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tRTI  Return from Interrupt (P and PC from Stack).\tP = Memory[0x%02X] %d | PC = 0x%04X | SP: 0x%02X\n", opc_string, SP_Address, P, PC, SP)
		fmt.Println(dbg_show_message)
	}
}
