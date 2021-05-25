package CORE

// import	"os"
import (
	"fmt"
)

// ROR  Rotate One Bit Right (Memory or Accumulator)
//
//      C -> [76543210] -> C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator      ROR A      6A    1     2
//      zeropage         ROR oper   66    2     5
//      zeropage,X       ROR oper,X 76    2     6
//      absolute         ROR oper   6E    3     6
//      absolute,X       ROR oper,X 7E    3     7

// Move each of the bits in either A or M one place to the right.
// Bit 7 is filled with the current value of the carry flag whilst the old bit 0 becomes the new carry flag value.

// ------------------------------------ Accumulator ------------------------------------ //

func opc_ROR_A(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Keep original Accumulator value for debug
		original_A := A
		original_carry := P[0]

		// Keep the original bit 0 from Accumulator to be used as new Carry
		new_Carry := A & 0x01

		// Shift Right Accumulator
		A = A >> 1

		// Bit 7 is filled with the current value of the carry flag
		A += (P[0] << 7)

		// The old bit 0 becomes the new carry flag value
		P[0] = new_Carry

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Accumulator]\tROR  Rotate One Bit Right.\tA(%d) Roll Right 1 bit\t(%d) + Current Carry(%d) as new bit 7.\tA = %d\n", opcode, original_A, original_A>>1, original_carry, A)
			fmt.Println(dbg_show_message)
		}

		flags_N(A)
		flags_Z(A)

		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", original_carry, P[0])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}

// --------------------------------------- Memory -------------------------------------- //

func opc_ROR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Keep original Accumulator value for debug
		original_MemValue := Memory[memAddr]
		original_carry := P[0]

		// Keep the original bit 0 from Accumulator to be used as new Carry
		new_Carry := Memory[memAddr] & 0x01

		// Shift Right Memory Value
		Memory[memAddr] = Memory[memAddr] >> 1

		// Bit 7 is filled with the current value of the carry flag
		Memory[memAddr] += (P[0] << 7)

		// The old bit 0 becomes the new carry flag value
		P[0] = new_Carry

		if Debug {

			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROR  Rotate One Bit Right.\tMemory[%02d](%d) Roll Right 1 bit\t(%d) + Current Carry(%d) as new bit 7.\tA = %d\n", opcode, Memory[PC+1], mode, memAddr, original_MemValue, original_MemValue>>1, original_carry, Memory[memAddr])
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tROR  Rotate One Bit Right.\tMemory[%02d](%d) Roll Right 1 bit\t(%d) + Current Carry(%d) as new bit 7.\tA = %d\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, original_MemValue, original_MemValue>>1, original_carry, Memory[memAddr])
			}
			fmt.Println(dbg_show_message)

		}

		flags_N(Memory[memAddr])
		flags_Z(Memory[memAddr])

		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", original_carry, P[0])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

	}

}
