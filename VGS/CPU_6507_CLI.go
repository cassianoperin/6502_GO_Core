package VGS

import "fmt"

// CLI  Clear Interrupt Disable Bit
//
//      0 -> I                           N Z C I D V
//                                       - - - 0 - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       CLI           58     1     2

func opc_CLI(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		P[2] = 0

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tCLI  Clear Interrupt Disable Bit.\tP[2]=%d\n", opcode, P[2])
			fmt.Println(dbg_show_message)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
