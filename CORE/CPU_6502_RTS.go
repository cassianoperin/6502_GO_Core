package CORE

import (
	"fmt"
)

// RTS  Return from Subroutine
//
//      pull PC, PC+1 -> PC              N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       RTS           60    1     6

func opc_RTS(bytes uint16, opc_cycles byte) {

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

		PC = uint16(Memory[SP_Address+2])<<8 | uint16(Memory[SP_Address+1])

		// UNDOCUMENTED // Clear the addresses retrieved from the stack
		Memory[SP_Address+1] = 0
		Memory[SP_Address+2] = 0

		// Update the Stack Pointer (Increase the two values retrieved)
		SP += 2

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tRTS  Return from Subroutine.\tPC = %04X (+ 1 RTS instruction byte).\n", opcode, PC)
			fmt.Println(dbg_show_message)
		}

		// Increment PC
		PC += bytes

		// // TEST THE MODIFICATIONS FROM SP_Address
		// fmt.Println("RTS - Validate the SP_Address in 6502 mode. Exiting.")
		// os.Exit(2)

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
