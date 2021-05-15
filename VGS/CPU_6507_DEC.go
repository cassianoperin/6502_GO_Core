package VGS

import "fmt"

// DEC  Decrement Memory by One
//
//      M - 1 -> M                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      DEC oper      C6    2     5
//      zeropage,X    DEC oper,X    D6    2     6
//      absolute      DEC oper      CE    3     6
//      absolute,X    DEC oper,X    DE    3     7

func opc_DEC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		if Debug {

			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tDEC  Decrement Memory by One.\tMemory[%02X](%d) - 1:\t%d\n", opcode, Memory[PC+1], mode, memAddr, Memory[memAddr], Memory[memAddr]-1)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tDEC  Decrement Memory by One.\tMemory[%02X](%d) - 1:\t%d\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Memory[memAddr], Memory[memAddr]-1)
			}
			fmt.Println(dbg_show_message)

		}

		Memory[memAddr] -= 1

		flags_Z(Memory[memAddr])
		flags_N(Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
