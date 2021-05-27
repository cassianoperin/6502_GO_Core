package CORE

import (
	"fmt"
)

// CMP  Compare Memory with Accumulator
//
//      A - M                          N Z C I D V
//                                     + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CMP #oper     C9    2     2
//      zeropage      CMP oper      C5    2     3
//      zeropage,X    CMP oper,X    D5    2     4
//      absolute      CMP oper      CD    3     4
//      absolute,X    CMP oper,X    DD    3     4*
//      absolute,Y    CMP oper,Y    D9    3     4*
//      (indirect,X)  CMP (oper,X)  C1    2     6
//      (indirect),Y  CMP (oper),Y  D1    2     5*

func opc_CMP(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles+opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles+opc_cycle_extra {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		var tmp byte

		tmp = A - Memory[memAddr]

		if Debug {

			if tmp == 0 {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+2], Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				}

			} else {

				if bytes == 2 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				} else if bytes == 3 {
					dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+2], Memory[PC+1], mode, A, memAddr, Memory[memAddr], tmp)
					fmt.Println(dbg_show_message)
				}

			}

		}

		flags_Z(tmp)
		flags_N(tmp)
		// Set if A >= M
		flags_C_CPX_CPY_CMP(A, Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
