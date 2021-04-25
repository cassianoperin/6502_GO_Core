package VGS

import (
	"fmt"
	"os"
)

// PLA  Pull Accumulator from Stack
//
//      pull A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PLA           68    1     4

func opc_PLA(bytes uint16, opc_cycles byte) {

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
			fmt.Printf("%d PLA TEST!", SP_Address)
			os.Exit(2)

			// 6502/6507 interpreter mode
		} else {
			// Stack is a 256-byte array whose location is hardcoded at page $01 ($0100-$01FF)
			SP_Address = uint(SP+1) + 256
		}

		A = Memory[SP_Address]

		// Not documented, clean the value on the stack after pull it to accumulator
		//Memory[SP_Address] = 0

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tPLA  Pull Accumulator from Stack.\tA = Memory[%02X] (%d) | SP++\n", opcode, SP_Address, A)
			fmt.Println(dbg_show_message)
		}

		flags_N(A)
		flags_Z(A)

		SP++

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
