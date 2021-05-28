package CORE

import "fmt"

// DEC  Decrement Memory by One
//
//      M - 1 -> M                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      DEC oper      C6    2     5
//      zeropage,X    DEC oper,X    D6    2     6
//      absolute      DEC oper      CE    3     6
//      absolute,X    DEC oper,X    DE    3     7

func opc_DEC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Print Opcode Debug Message
		opc_DEC_DebugMsg(bytes, mode, memAddr)

		Memory[memAddr] -= 1

		flags_Z(Memory[memAddr])
		flags_N(Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}

}

func opc_DEC_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tDEC  Decrement Memory by One.\tMemory[0x%02X](%d) - 1:\t%d\n", opc_string, mode, memAddr, Memory[memAddr], Memory[memAddr]-1)
		fmt.Println(dbg_show_message)
	}
}
