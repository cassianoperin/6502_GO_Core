package CORE

import (
	"fmt"
)

// AND  AND Memory with Accumulator
//
//      A AND M -> A                     N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     AND #oper     29    2     2
//      zeropage      AND oper    	25    2   	3
//      zeropage,X    AND oper,X    35    2     4
//      absolute      AND oper      2D    3     4
//      absolute,X    AND oper,X    3D    3     4*
//      absolute,Y    AND oper,Y    39    3     4*
//      (indirect,X)  AND (oper,X)  21    2     6
//      (indirect),Y  AND (oper),Y  31    2     5*

func opc_AND(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Print internal opcode cycle
	debugInternalOpcCycleExtras(opc_cycles)

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles+opc_cycle_extra {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Print Opcode Debug Message
		opc_AND_DebugMsg(bytes, mode, memAddr)

		A = A & Memory[memAddr]

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_AND_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tAND  AND Memory with Accumulator.\tA = A(%d) & Memory[0x%02X](%d)\t(%d)\n", opc_string, mode, A, memAddr, Memory[memAddr], A&Memory[memAddr])
		fmt.Println(dbg_show_message)
	}
}
