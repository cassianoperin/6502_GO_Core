package CORE

import (
	"fmt"
)

// NOP  No Operation (Unofficial)
//
//      ---                              N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      NOP oper      64    2     3

func opc_U_NOP(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Print Opcode Debug Message
		opc_U_NOP_DebugMsg(bytes, mode)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}

}

func opc_U_NOP_DebugMsg(bytes uint16, mode string) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s] [Unnoficial!!!]\tNOP  No Operation, ignore next byte (%02X).\n", opc_string, mode, Memory[PC+1])
		fmt.Println(dbg_show_message)
	}
}
