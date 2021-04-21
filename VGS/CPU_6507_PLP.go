package VGS

import (
	"fmt"
	"os"
)

// Pull  Processor Status from Stack
//
//      pull SR                          N Z C I D V
//                                       from stack
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PLP           28    1     4

func opc_PLP(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		var SP_Address uint

		// Atari 2600 interpreter mode
		if CPU_MODE == 0 {
			SP_Address = uint(SP + 1)

			// Test
			fmt.Printf("%d PLP TEST!", SP_Address)
			os.Exit(2)

			// 6502/6507 interpreter mode
		} else {
			// Stack is a 256-byte array whose location is hardcoded at page $01 ($0100-$01FF)
			SP_Address = uint(SP+1) + 256
		}

		// Turn the stack value into the processor status
		for i := 0; i < len(P); i++ {
			P[i] = (Memory[SP_Address] >> i) & 0x01
		}

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tPLP  Processor Status from Stack.\tP = Memory[%02X] %d | SP++\n", opcode, SP_Address, P)
			fmt.Println(dbg_show_message)
		}

		SP++

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
