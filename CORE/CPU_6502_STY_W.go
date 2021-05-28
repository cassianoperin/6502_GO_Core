package CORE

import "fmt"

// STY  Store Index Y in Memory (zeropage)
//
//      Y -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STY oper      84    2     3
//      zeropage,X    STY oper,X    94    2     4
//      absolute      STY oper      8C    3     4

func opc_STY(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Update Memory[memAddr] with value of Y and notify TIA about the update
		memUpdate(memAddr, Y)

		// Print Opcode Debug Message
		opc_STY_DebugMsg(bytes, mode, memAddr)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_STY_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tSTY  Store Index Y in Memory.\tMemory[0x%02X] = Y (%d)\n", opc_string, mode, memAddr, Y)
		fmt.Println(dbg_show_message)
	}
}
