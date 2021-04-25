package VGS

import "fmt"

// LDY  Load Index Y with Memory (immediate)
//
//      M -> Y                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDY #oper     A0    2     2
//      zeropage      LDY oper      A4    2     3
//      zeropage,X    LDY oper,X    B4    2     4
//      absolute      LDY oper      AC    3     4
//      absolute,X    LDY oper,X    BC    3     4*

func opc_LDY(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Check for extra cycles (*) in the first opcode cycle
	if opc_cycle_count == 1 {
		if opcode == 0xBC {
			// Add 1 to cycles if page boundery is crossed
			if MemPageBoundary(memAddr, PC) {
				opc_cycle_extra = 1
			}
		}
	}

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles+opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles+opc_cycle_extra {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		Y = Memory[memAddr]

		if Debug {
			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDY  Index Y with Memory.\tY = Memory[%02X] (%d)\n", opcode, Memory[PC+1], mode, memAddr, Y)
				fmt.Println(dbg_show_message)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tLDY  Index Y with Memory.\tY = Memory[%02X] (%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Y)
				fmt.Println(dbg_show_message)
			}
		}

		flags_Z(Y)
		flags_N(Y)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
