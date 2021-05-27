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

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles+opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles+opc_cycle_extra {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		if Debug {

			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tAND  AND Memory with Accumulator.\tA = A(%d) & Memory[%02X](%d)\t(%d)\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], A&Memory[memAddr])
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tAND  AND Memory with Accumulator.\tA = A(%d) & Memory[%02X](%d)\t(%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, A, memAddr, Memory[memAddr], A&Memory[memAddr])
			}
			fmt.Println(dbg_show_message)

		}

		A = A & Memory[memAddr]

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
