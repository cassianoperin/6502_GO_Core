package CORE

import "fmt"

// ROL  Rotate One Bit Left (Memory or Accumulator)
//
//      C <- [76543210] <- C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   ROL A         2A    1     2
//      zeropage      ROL oper      26    2     5
//      zeropage,X    ROL oper,X    36    2     6
//      absolute      ROL oper      2E    3     6
//      absolute,X    ROL oper,X    3E    3     7

//Move each of the bits in either A or M one place to the left.
//Bit 0 is filled with the current value of the carry flag whilst the old bit 7 becomes the new carry flag value.

// ------------------------------------ Accumulator ------------------------------------ //

func opc_ROL_A(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Original Carry Value
		carry_orig := P[0]

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Accumulator]\tROL  Rotate One Bit Left.\tA(%d) Roll Left 1 bit + carry(%d)\t: %d\n", opcode, A, P[0], (A<<1)+carry_orig)
			fmt.Println(dbg_show_message)
		}

		// Calculate the original bit7 and save it as the new Carry
		P[0] = A & 0x80 >> 7

		// Shift left the byte and put the original bit7 value in bit 1 to make the complete ROL
		A = (A << 1) + carry_orig

		flags_N(A)
		flags_Z(A)
		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", carry_orig, P[0])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

// --------------------------------------- Memory -------------------------------------- //

func opc_ROL(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Original Carry Value
		carry_orig := P[0]

		if Debug {

			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROL  Rotate One Bit Left.\tMemory[%d](%d) Roll Left 1 bit + Carry(%d)\t(%d)\n", opcode, Memory[PC+1], mode, memAddr, Memory[memAddr], carry_orig, (Memory[memAddr]<<1)+carry_orig)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tROL  Rotate One Bit Left.\tMemory[%d](%d) Roll Left 1 bit + Carry(%d)\t(%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Memory[memAddr], carry_orig, (Memory[memAddr]<<1)+carry_orig)
			}
			fmt.Println(dbg_show_message)

		}

		// Calculate the original bit7 and save it as the new Carry
		P[0] = Memory[memAddr] & 0x80 >> 7

		// Shift left the byte and put the original bit7 value in bit 1 to make the complete ROL
		Memory[memAddr] = (Memory[memAddr] << 1) + carry_orig

		flags_N(Memory[memAddr])
		flags_Z(Memory[memAddr])
		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", carry_orig, P[0])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
