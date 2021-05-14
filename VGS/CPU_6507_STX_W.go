package VGS

import "fmt"

// STX  Store Index X in Memory
//
//      X -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STX oper      86    2     3
//      zeropage,Y    STX oper,Y    96    2     4
//      absolute      STX oper      8E    3     4

func opc_STX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Update Memory[memAddr] with value of X and notify TIA about the update
		memUpdate(memAddr, X)

		if Debug {

			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTX  Store Index X in Memory.\tMemory[%02X] = X (%d)\n", opcode, Memory[PC+1], mode, memAddr, X)
				fmt.Println(dbg_show_message)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSTX  Store Index X in Memory.\tMemory[%02X] = X (%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, X)
				fmt.Println(dbg_show_message)
			}

		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
