package CORE

import "fmt"

// CPY  Compare Memory and Index Y
//
//      Y - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CPY #oper     C0    2     2
//      zeropage      CPY oper      C4    2     3
//      absolute      CPY oper      CC    3     4

func opc_CPY(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		tmp := Y - Memory[memAddr]

		// Print Opcode Debug Message
		opc_CPY_DebugMsg(bytes, tmp, mode, memAddr)

		// Set if Y = M
		flags_Z(tmp)
		// Set if bit 7 of the result is set
		flags_N(tmp)
		// Set if Y >= M
		flags_C_CPX_CPY_CMP(Y, Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}

}

func opc_CPY_DebugMsg(bytes uint16, tmp byte, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if tmp == 0 {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCPY  Compare Memory and Index Y.\tY(%d) - Memory[0x%02X](%d) = (%d) EQUAL\n", opc_string, mode, Y, PC+1, Memory[memAddr], tmp)
		} else {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCPY  Compare Memory and Index Y.\tY(%d) - Memory[0x%02X](%d) = (%d) NOT EQUAL\n", opc_string, mode, Y, PC+1, Memory[memAddr], tmp)
		}
		fmt.Println(dbg_show_message)
	}
}
