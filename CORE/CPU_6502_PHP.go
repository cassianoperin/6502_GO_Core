package CORE

import (
	"fmt"
)

//      PHP  Push Processor Status on Stack
//
//      push SR                          N Z C I D V
//                                       - - - - - -
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PHP           08    1     3

func opc_PHP(bytes uint16, opc_cycles byte) {

	var tmp_P byte

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		var SP_Address uint16 = uint16(SP) + 256 // 6502 handle Stack at the end of first memory page

		// Put processor Status (P) on stack
		for i := 7; i >= 0; i-- {

			// The B Flag, for PHP or BRK, P[4] and P[5] will be always 1
			if i == 4 || i == 5 {
				tmp_P = (tmp_P << 1) + 1
			} else {
				tmp_P = (tmp_P << 1) + P[i]
			}

		}

		Memory[SP_Address] = tmp_P

		// Print Opcode Debug Message
		opc_PHP_DebugMsg(bytes, SP_Address, tmp_P)

		SP--

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}

}

func opc_PHP_DebugMsg(bytes uint16, SP_Address uint16, tmp_P byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tPHP  Push Processor Status on Stack.\tMemory[0x%02X] = Processor Status %08b | SP--\n", opc_string, SP_Address, tmp_P)
		fmt.Println(dbg_show_message)
	}
}
