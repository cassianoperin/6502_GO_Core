package CORE

import "fmt"

// LSR  Shift One Bit Right (Memory or Accumulator)
//
//      0 -> [76543210] -> C             N Z C I D V
//                                       0 + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   LSR A         4A    1     2
//      zeropage      LSR oper      46    2     5
//      zeropage,X    LSR oper,X    56    2     6
//      absolute      LSR oper      4E    3     6
//      absolute,X    LSR oper,X    5E    3     7

// ------------------------------------ Accumulator ------------------------------------ //

func opc_LSR_A(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Save the original Carry value
		carry_orig := P[0]

		// Least significant bit turns into the new Carry
		P[0] = A & 0x01

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Accumulator]\tLSR  Shift One Bit Right.\tA = A(%d) Shift Right 1 bit\t(%d)\n", opcode, A, A>>1)
			fmt.Println(dbg_show_message)
		}

		A = A >> 1

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

func opc_LSR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Save the original Carry value
		carry_orig := P[0]

		// Least significant bit turns into the new Carry
		P[0] = Memory[memAddr] & 0x01

		if Debug {

			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLSR  Shift One Bit Right.\tMemory[%d]: (%d) Shift Right 1 bit\t(%d)\n", opcode, Memory[PC+1], mode, memAddr, Memory[memAddr], Memory[memAddr]>>1)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tLSR  Shift One Bit Right.\tMemory[%d]: (%d) Shift Right 1 bit\t(%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Memory[memAddr], Memory[memAddr]>>1)
			}
			fmt.Println(dbg_show_message)

		}

		Memory[memAddr] = Memory[memAddr] >> 1

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
