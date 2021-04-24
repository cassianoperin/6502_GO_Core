package VGS

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

		// Push PC+2 (will be increased in 1 in RTS to match the next address (3 bytes operation))
		// Store the first byte into the Stack
		// fmt.Printf("PC: %02X\n", PC)

		Memory[SP_Address] = byte((PC + 2) >> 8)
		// fmt.Printf("FF %02X\n", Memory[SP_Address])
		SP--
		SP_Address--
		// Store the second byte into the Stack
		Memory[SP_Address] = byte((PC + 2) & 0xFF)
		SP_Address--
		SP--
		// fmt.Printf("FE %02X\n", Memory[SP_Address])

		// fmt.Printf("\nPC+3: %02X",PC+3)
		// fmt.Printf("\nF0: %02X",(PC+3) >> 8)
		// fmt.Printf("\n42: %02X",(PC+3) & 0xFF)

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tJSR  Jump to New Location Saving Return Address.\tPC = Memory[%02X]\t|\t Stack[%02X] = %02X\t Stack[%02X] = %02X\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, SP_Address+2, Memory[SP_Address+2], SP_Address+1, Memory[SP_Address+1])
			fmt.Println(dbg_show_message)
		}

		// TEST THE MODIFICATIONS FROM SP_Address
		// fmt.Println("JSR - Validate the SP_Address in 6502 mode. Exiting.")
		// os.Exit(2)

		// Update PC
		PC = memAddr

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
