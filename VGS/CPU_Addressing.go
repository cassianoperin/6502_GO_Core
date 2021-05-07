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

// Zeropage,Y
func addr_mode_ZeropageY(offset uint16) (uint16, string) {

	value := Memory[Memory[offset]+Y]
	memAddr := Memory[offset] + Y
	mode := "Zeropage,Y"

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

	// PAUSE HERE TO FIX THE 6502 BUG WHEN THE ADDRESS is 0xFF
	// https://www.reddit.com/r/EmuDev/comments/fi29ah/6502_jump_indirect_error/
	//It's a bug in the 6502 that wraps around the LSB without incrementing the MSB.
	//So instead of reading address from 0x02FF-0x0300 you should be looking at 0x02FF-0x0200.
	// The A900 printed in the log is the value at 0x02FF-0x0300 which is not what's actually being used.
	if Memory[offset+1] == 0xFF || Memory[offset] == 0xFF {
		fmt.Printf("Controled Exit on Indirect Memory mode to correct a bug in 6502. Mem1: %02X Mem2: %02X. Exiting", Memory[offset+1], Memory[offset])
		os.Exit(2)
	}

	// First format the destination address
	memAddr := uint16(Memory[offset+1])<<8 | uint16(Memory[offset])
	// Get the value in the memory of this address (Indirect)
	memAddr = uint16(Memory[memAddr+1])<<8 | uint16(Memory[memAddr])
	// value := Memory[memAddr]
	mode := "Indirect"

	if Debug {
		fmt.Printf("\t%s addressing mode.\t\tAddress in $%04X points to Memory[%04X]\n", mode, (uint16(Memory[offset+1])<<8 | uint16(Memory[offset])), memAddr)
	}
	return memAddr, mode
}

// Indirect,Y
func addr_mode_IndirectY(offset uint16) (uint16, string) {

	var (
		indirect_addr, LSB, MSB, carry byte
		LSB_tmp                        uint16
	)

	// Base indirect address
	indirect_addr = Memory[offset]

	// In (Indirect),Y mode, its necessary to sum the memory inside the indirect address + Y and keep the carry if exists to use in MSB
	LSB_tmp = uint16(Memory[indirect_addr]) + uint16(Y)

	// Keep the bit 9 as the carry for MSB
	carry = byte(LSB_tmp >> 8)

	// Temporary test
	if carry == 1 {
		fmt.Printf("\n\nIndirect,Y mode, proposital exit: Y+indirect addr > 255, carry on. Validate the sum!!!\n\n")
		// os.Exit(2)
		Pause = true
	}

	// Store only the first 8 bits as LSB (the 9th is on carry)
	LSB = byte(LSB_tmp & 0x00FF)
	// Most significant bit will be memory inside the next address after indirect_add + Carry from LSB (if exist)
	MSB = Memory[indirect_addr+1+carry]

	memAddr := uint16(MSB)<<8 | uint16(LSB)
	value := Memory[memAddr]
	mode := "Indirect,Y"

	if Debug {
		fmt.Printf("\t%s addressing mode. Indirect Addr: 0x%02X\tLSB: (Memory[indirect_addr]:0x%02X + Y:(0x%02X)) = 0x%04X & 00FF = 0x%02X and carry: %d\tMSB: (Memory[indirect_addr(%02X) + 1 + carry(%d)]): 0x%02X\tMemory[%04X]\t\tValue obtained: %02X\n", mode, indirect_addr, Memory[indirect_addr], Y, LSB_tmp, LSB, carry, indirect_addr, carry, MSB, memAddr, value)
	}

	return memAddr, mode
}

// Indirect,X
func addr_mode_IndirectX(offset uint16) (uint16, string) {

	var (
		indirect_addr, LSB, MSB byte
	)

	// Base indirect address
	indirect_addr = Memory[offset]

	// In (Indirect,X) mode, its necessary to sum the address pointed on indirect address + X, ignoring the carry if exists
	// Store only the first 8 bits as LSB, ignoring Carry (byte sum will do it itself rotating the number if greater than 255)
	LSB = indirect_addr + X
	MSB = LSB + 0x01 // Next byte

	memAddr := uint16(Memory[MSB])<<8 | uint16(Memory[LSB])
	value := Memory[memAddr]
	mode := "Indirect,X"

	if Debug {
		fmt.Printf("\t%s addressing mode. Indirect Addr: 0x%02X\tLSB: (indirect_addr:0x%02X + X:(0x%02X) = 0x%02X\tMSB: Address of LSB + 0x01: 0x%02X\tMemory[%04X]\t\tValue obtained: %02X\n", mode, indirect_addr, indirect_addr, Y, LSB, MSB, memAddr, value)
	}

	return memAddr, mode
}
