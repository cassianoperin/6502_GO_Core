package CORE

import (
	"fmt"
)

// JSR  Jump to New Location Saving Return Address
//
//      push (PC+2) to Stack,            N Z C I D V
//      (PC+1) -> PCL                    - - - - - -
//      (PC+2) -> PCH
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      absolute      JSR oper      20    3     6

func opc_JSR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		Memory[SP_Address] = byte((PC + 2) >> 8)
		SP--
		SP_Address--
		// Store the second byte into the Stack
		Memory[SP_Address] = byte((PC + 2) & 0xFF)
		SP_Address--
		SP--

		// Print Opcode Debug Message
		opc_JSR_DebugMsg(bytes, mode, memAddr, SP_Address)

		// Update PC
		PC = memAddr

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

func opc_JSR_DebugMsg(bytes uint16, mode string, memAddr uint16, SP_Address uint) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tJSR  Jump to New Location Saving Return Address.\tPC = Memory[0x%02X]\t|\t Stack[0x%02X] = %02X\t Stack[0x%02X] = 0x%02X\n", opc_string, mode, memAddr, SP_Address+2, Memory[SP_Address+2], SP_Address+1, Memory[SP_Address+1])
		fmt.Println(dbg_show_message)
	}
}
