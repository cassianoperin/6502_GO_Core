package CORE

import (
	"fmt"
)

// BRK  Force Break
//
//      interrupt,                       N Z C I D V
//      push PC+2, push SR               - - - 1 - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       BRK           00    1     7

// Order
// store PC(hi)
// store PC(lo)
// store P
// fetch PC(lo) from $FFFE
// fetch PC(hi) from $FFFF

func opc_BRK(bytes uint16, opc_cycles byte) {

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// ---------- Store PC ---------- //

		var SP_Address uint

		// Atari 2600 interpreter mode
		if CPU_MODE == 0 {
			SP_Address = uint(SP)

			// 6502/6507 interpreter mode
		} else {
			// Stack is a 256-byte array whose location is hardcoded at page $01 ($0100-$01FF)
			SP_Address = uint(SP) + 256
		}

		// Push PC+2 (PC(hi))
		Memory[SP_Address] = byte((PC + 2) >> 8)
		SP--
		SP_Address--

		// Push PC+1 (PC(lo))
		Memory[SP_Address] = byte((PC + 2) & 0xFF)
		SP_Address--
		SP--

		// ---------- Store P ----------- //

		var tmp_P byte

		// Put processor Status (P) on stack
		for i := 7; i >= 0; i-- {

			// The B Flag, for PHP or BRK, P[4] and P[5] will be always 1
			if i == 4 || i == 5 {
				tmp_P = (tmp_P << 1) + 1
			} else {
				tmp_P = (tmp_P << 1) + P[i]
			}

		}

		// Push Processor Status (P) to Stack
		Memory[SP_Address] = tmp_P
		SP_Address--
		SP--

		// ---------- Fetch PC ---------- //

		// Read the Opcode from PC+1 and PC bytes (Little Endian)
		PC = uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE])

		// ------------ Flags ----------- //

		// IRQ Disabled
		P[2] = 1

		// The B Flag, for PHP or BRK, P[4] and P[5] will be always 1
		P[4] = 1
		// P[5] = 1

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tBRK  Force Break.\tPush PC and P to Stack: Mem[%02X] = %02X ,Mem[%02X] = %02X, Mem[%02X] = %02X(%08b)\t\tNew PC = %04X(BRK/Interrupt)\n", opcode, SP_Address+3, Memory[SP_Address+3], SP_Address+2, Memory[SP_Address+2], SP_Address+1, Memory[SP_Address+1], Memory[SP_Address+1], uint16(Memory[0xFFFF])<<8|uint16(Memory[0xFFFE]))
			println(dbg_show_message)
		}

	}

}
