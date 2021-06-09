package CORE

import (
	"fmt"
)

// ---------------------------- Library Function ---------------------------- //

// Memory Page Boundary cross detection
func MemPageBoundary(original_addr, new_addr uint16) byte {

	var extra_cycle byte = 0

	// Page Boundary Cross detected
	if original_addr>>8 != new_addr>>8 { // Get the High byte only to compare

		extra_cycle = 1

		if Debug {
			fmt.Printf("\tMemory Page Boundary Cross detected! Add 1 cycle.\tPC High byte: %02X\tBranch High byte: %02X\n", original_addr>>8, new_addr>>8)
		}

		// NO Page Boundary Cross detected
	} else {

		extra_cycle = 0

		if Debug {
			fmt.Printf("\tNo Memory Page Boundary Cross detected.\tPC High byte: %02X\tBranch High byte: %02X\n", original_addr>>8, new_addr>>8)
		}
	}

	return extra_cycle
}

// Decode Two's Complement
func DecodeTwoComplement(num byte) int8 {

	var sum int8 = 0

	for i := 0; i < 8; i++ {
		// Sum each bit and sum the value of the bit power of i (<<i)
		sum += (int8(num) >> i & 0x01) << i
	}

	return sum
}

// Decode opcode for debug messages
func debug_decode_opc(bytes uint16) string {

	var opc_string string

	// Decode opcode and operators
	for i := 0; i < int(bytes); i++ {
		if i == 1 {
			opc_string += fmt.Sprintf(" %02X", Memory[PC+uint16(i)])
		} else {
			opc_string += fmt.Sprintf("%02X", Memory[PC+uint16(i)])
		}
	}

	// Insert number of bytes into the string
	if bytes == 1 {
		opc_string += " [1 byte]"
	} else if bytes == 2 {
		opc_string += " [2 bytes]"
	} else {
		opc_string += " [3 bytes]"
	}

	return opc_string
}

// Print internal opcode cycle in debug mode
func debugInternalOpcCycle(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", cycle, opc_cycle_count, opc_cycles)
	}
}

// Print internal opcode cycle in debug mode - instructions with extra cycle
func debugInternalOpcCycleExtras(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", cycle, opc_cycle_count, opc_cycles+opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}
}

// Print internal opcode cycle in debug mode - Branches
func debugInternalOpcCycleBranch(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + 1 cycle for branch + %d extra cycles for branch in different page)\n", cycle, opc_cycle_count, opc_cycles+opc_cycle_extra+1, opc_cycles, opc_cycle_extra)
	}
}

// Reset the internal opcode cycle counters
func resetIntOpcCycleCounters() {
	// Reset Opcode Cycle counter
	opc_cycle_count = 1

	// Reset Opcode Extra Cycle counter
	opc_cycle_extra = 0

	// Update IPS
	IPS++
}

// Data Bus - READ from Memory Operations
func dataBUS_Read(memAddr uint16) byte {
	data_value := Memory[memAddr]

	return data_value
}

// Data Bus - WRITE to Memory Operations
func dataBUS_Write(memAddr uint16, data_value byte) byte {

	Memory[memAddr] = data_value

	return data_value
}
