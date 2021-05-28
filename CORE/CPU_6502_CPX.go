package CORE

import "fmt"

// CPX  Compare Memory and Index X
//
//      X - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CPX #oper     E0    2     2
//      zeropage      CPX oper    	E4    2	    3
//      absolute      CPX oper      EC    3     4

func opc_CPX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		tmp := X - Memory[memAddr]

		// Print Opcode Debug Message
		opc_CPX_DebugMsg(bytes, tmp, mode, memAddr)

		// Set if X = M
		flags_Z(tmp)
		// Set if bit 7 of the result is set
		flags_N(tmp)
		// Set if X >= M
		flags_C_CPX_CPY_CMP(X, Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

func opc_CPX_DebugMsg(bytes uint16, tmp byte, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if tmp == 0 {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[0x%02X](%d) = (%d) EQUAL\n", opc_string, mode, X, PC+1, Memory[memAddr], tmp)
		} else {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[0x%02X](%d) = (%d) NOT EQUAL\n", opc_string, mode, X, PC+1, Memory[memAddr], tmp)
		}
		fmt.Println(dbg_show_message)
	}
}
