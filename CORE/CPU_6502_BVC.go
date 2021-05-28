package CORE

import "fmt"

// BVC  Branch on Overflow Clear
//
//      branch on V = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BVC oper      50    2     2**

func opc_BVC(value int8, bytes uint16, opc_cycles byte) { // value is SIGNED

	if P[6] == 0 { // If Overflow is clear

		// Show current opcode cycle
		if Debug {
			fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + 1 cycle for branch + %d extra cycles for branch in different page)\n", counter_F_Cycle, opc_cycle_count, opc_cycles+opc_cycle_extra+1, opc_cycles, opc_cycle_extra)
		}

		// Just increment the Opcode cycle Counter
		if opc_cycle_count < opc_cycles+1+opc_cycle_extra {
			opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			// Print Opcode Debug Message
			opc_BVC_DebugMsg(bytes, value)

			// PC + the number of bytes to jump on Overflow clear
			PC += uint16(value)

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1

			// Reset Opcode Extra Cycle counter
			opc_cycle_extra = 0
		}

	} else { // If Overflow is set

		// Show current opcode cycle
		if Debug {
			fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
		}

		// Just increment the Opcode cycle Counter
		if opc_cycle_count < opc_cycles {
			opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			// Print Opcode Debug Message
			opc_BVC_DebugMsg(bytes, value)

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1
		}
	}
}

func opc_BVC_DebugMsg(bytes uint16, value int8) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if P[6] == 0 { // If Overflow is clear
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Relative]\tBVC  Branch on Overflow Clear.\tOverflow EQUAL 0, JUMP TO 0x%04X\n", opc_string, PC+2+uint16(value))
		} else { // If Overflow is set
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s\tBVC  Branch on Overflow Clear.\tOverflow NOT EQUAL 0, PC+2\n", opc_string)
		}
		fmt.Println(dbg_show_message)
	}
}
