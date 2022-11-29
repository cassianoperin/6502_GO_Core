package CPU_6502

import "fmt"

// INX  Increment Index X by One
//
//      X + 1 -> X                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       INX           E8    1     2

func opc_INX(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		X++

		// Print Opcode Debug Message
		opc_INX_DebugMsg(bytes)

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_INX_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tINX  Increment Index X by One (0x%02X)\n", opc_string, X)
		fmt.Println(dbg_show_message)
	}
}