package VGS

import (
	"fmt"
)

//      PHP  Push Processor Status on Stack
//
//      push SR                          N Z C I D V
//                                       - - - - - -
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PHP           08    1     3
func opc_PHP(bytes uint16, opc_cycles byte) {

	var tmp_P byte

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
			SP_Address = uint(SP)

			// 6502/6507 interpreter mode
		} else {
			// Stack is a 256-byte array whose location is hardcoded at page $01 ($0100-$01FF)
			SP_Address = uint(SP) + 256
		}

		// Put processor Status (P) on stack
		for i := 7; i >= 0; i-- {

			// The B Flag, for PHP or BRK, P[4] and P[5] will be always 1
			if i == 4 || i == 5 {
				tmp_P = (tmp_P << 1) + 1
			} else {
				tmp_P = (tmp_P << 1) + P[i]
			}

		}

		Memory[SP_Address] = tmp_P

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tPHP  Push Processor Status on Stack.\tMemory[%02X] = Processor Status %08b | SP--\n", opcode, SP_Address, tmp_P)
			fmt.Println(dbg_show_message)
		}

		SP--

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
