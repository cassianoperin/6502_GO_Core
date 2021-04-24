package VGS

import (
	"fmt"
	"os"
)

// Relative
func addr_mode_Relative(offset uint16) int8 {

	// Branches needs the Two Complement of the offset value
	value := DecodeTwoComplement(Memory[offset])
	memAddr := offset
	mode := "Relative"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d (SIGNED value)\n", mode, memAddr, value)
	}

	return value
}

// Zeropage
func addr_mode_Zeropage(offset uint16) (uint16, string) {

	value := Memory[Memory[offset]]
	memAddr := Memory[offset]
	mode := "Zeropage"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d\n", mode, memAddr, value)
	}

	return uint16(memAddr), mode
}

// Zeropage,X
func addr_mode_ZeropageX(offset uint16) (uint16, string) {

	value := Memory[Memory[offset]+X]
	memAddr := Memory[offset] + X
	mode := "Zeropage,X"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d\n", mode, memAddr, value)
	}

	return uint16(memAddr), mode
}

// Immediate
func addr_mode_Immediate(offset uint16) (uint16, string) {

	value := Memory[offset]
	memAddr := offset
	mode := "Immediate"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d\n", mode, memAddr, value)
	}

	return memAddr, mode

}

// Absolute
func addr_mode_Absolute(offset uint16) (uint16, string) {

	memAddr := uint16(Memory[offset+1])<<8 | uint16(Memory[offset])
	value := Memory[memAddr]
	mode := "Absolute"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\t\tValue obtained: %d\n", mode, memAddr, value)
	}
	return memAddr, mode
}

// Absolute,Y
func addr_mode_AbsoluteY(offset uint16) (uint16, string) {

	memAddr := (uint16(Memory[offset+1])<<8 | uint16(Memory[offset])) + uint16(Y)
	value := Memory[memAddr]
	mode := "Absolute,Y"

	if Debug {
		fmt.Printf("\t%s addressing mode.\t\tMemory[%02X]\t\tValue obtained: %d\n", mode, memAddr, value)
	}
	return memAddr, mode
}

// Absolute,X
func addr_mode_AbsoluteX(offset uint16) (uint16, string) {

	memAddr := (uint16(Memory[offset+1])<<8 | uint16(Memory[offset])) + uint16(X)
	value := Memory[memAddr]
	mode := "Absolute,X"

	if Debug {
		fmt.Printf("\t%s addressing mode.\t\tMemory[%02X]\t\tValue obtained: %d\n", mode, memAddr, value)
	}

	return memAddr, mode
}

// Indirect
func addr_mode_Indirect(offset uint16) (uint16, string) {

	// First format the destination address
	memAddr := (uint16(Memory[offset+1])<<8 | uint16(Memory[offset]))

	// PAUSE HERE TO FIX THE 6502 BUG WHEN THE ADDRESS is 0xFF
	// https://www.reddit.com/r/EmuDev/comments/fi29ah/6502_jump_indirect_error/
	//It's a bug in the 6502 that wraps around the LSB without incrementing the MSB.
	//So instead of reading address from 0x02FF-0x0300 you should be looking at 0x02FF-0x0200.
	// The A900 printed in the log is the value at 0x02FF-0x0300 which is not what's actually being used.
	if Memory[offset+1] == 0xFF || Memory[offset] == 0xFF {
		fmt.Printf("Controled Exit on Indirect Memory mode to correct a bug in 6502. Mem1: %02X Mem2: %02X. Exiting", Memory[offset+1], Memory[offset])
		os.Exit(2)
	}

	// Get the value in the memory of this address (Indirect)
	memAddr = uint16(Memory[memAddr+1])<<8 | uint16(Memory[memAddr])
	value := Memory[memAddr]
	mode := "Indirect"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%04X]\t\tValue obtained: %02X\n", mode, memAddr, value)
	}
	return memAddr, mode
}

// Indirect,Y
func addr_mode_IndirectY(offset uint16) (uint16, string) {

	memAddr := (uint16(Memory[Memory[offset+1]])<<8 | uint16(Memory[Memory[offset]])) + uint16(Y)
	value := Memory[memAddr]
	mode := "Indirect,Y"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%04X]\t\tValue obtained: %02X\n", mode, memAddr, value)
	}
	return memAddr, mode
}

// Indirect,X
func addr_mode_IndirectX(offset uint16) (uint16, string) {

	memAddr := (uint16(Memory[Memory[offset+1]])<<8 | uint16(Memory[Memory[offset]])) + uint16(X)
	value := Memory[memAddr]
	mode := "Indirect,X"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%04X]\t\tValue obtained: %02X\n", mode, memAddr, value)
	}
	return memAddr, mode
}

// // Accumulator
// func addr_mode_Accumulator(offset uint16) (uint16, string) {

// 	memAddr := (uint16(Memory[Memory[offset+1]])<<8 | uint16(Memory[Memory[offset]])) + uint16(X)
// 	value := Memory[memAddr]
// 	mode := "Accumulator"

// 	if Debug {
// 		fmt.Printf("\t%s addressing mode.\tMemory[%04X]\t\tValue obtained: %02X\n", mode, memAddr, value)
// 	}
// 	return memAddr, mode
// }
