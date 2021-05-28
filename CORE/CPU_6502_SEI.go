package CORE

import "fmt"

// SEI  Set Interrupt Disable Status
//
//      1 -> I                           N Z C I D V
//                                       - - - 1 - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       SEI           78    1     2

func opc_SEI(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		P[2] = 1

		// Print Opcode Debug Message
		opc_SEI_DebugMsg(bytes)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}
}

func opc_SEI_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tSEI  Set Interrupt Disable Status.\tP[2]=%d\n", opc_string, P[2])
		fmt.Println(dbg_show_message)
	}
}
