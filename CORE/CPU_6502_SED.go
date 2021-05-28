package CORE

import "fmt"

// SED  Set Decimal Flag
//
//     1 -> D                            N Z C I D V
//                                       - - - - 1 -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       SED           F8    1     2

func opc_SED(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		P[3] = 1

		// Print Opcode Debug Message
		opc_SED_DebugMsg(bytes)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

func opc_SED_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tSED   Set Decimal Flag.\tP[3]=1\n", opc_string)
		fmt.Println(dbg_show_message)
	}
}
