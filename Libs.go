package CPU_6502

import (
	"fmt"
	"log"
	"os"
)

// ---------------------------- Library Function ---------------------------- //

// Function used by ReadROM to avoid 'bytesread' return
func ReadContent(file *os.File, bytes_number int) []byte {

	bytes := make([]byte, bytes_number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

// Read ROM and write it to the RAM
func ReadROM(filename string) {

	var (
		fileInfo os.FileInfo
		err      error
	)

	// Get ROM info
	fileInfo, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loading ROM:", filename)
	romsize := fileInfo.Size()
	fmt.Printf("Size in bytes: %d\n", romsize)

	// Open ROM file, insert all bytes into memory
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Call ReadContent passing the total size of bytes
	data := ReadContent(file, int(romsize))
	// Print raw data
	//fmt.Printf("%d\n", data)
	//fmt.Printf("%X\n", data)

	// // 4KB roms
	// if romsize == 4096 {
	// 	// Load ROM to memory
	// 	for i := 0; i < len(data); i++ {
	// 		// F000 - FFFF // Cartridge ROM
	// 		VGS.Memory[0xF000+i] = data[i]
	// 	}
	// }

	// // 2KB roms (needs to duplicate it in memory)
	// if romsize == 2048 {
	// 	// Load ROM to memory
	// 	for i := 0; i < len(data); i++ {
	// 		// F000 - F7FF (2KB Cartridge ROM)
	// 		VGS.Memory[0xF000+i] = data[i]
	// 		// F800 - FFFF (2KB Mirror Cartridge ROM)
	// 		VGS.Memory[0xF800+i] = data[i]
	// 	}
	// }

	if romsize <= 65536 {
		// Load ROM to memory
		for i := 0; i < len(data); i++ {
			// F000 - F7FF (2KB Cartridge ROM)
			// Memory[i] = data[i]
			// F800 - FFFF (2KB Mirror Cartridge ROM)
			Memory[i] = data[i]
		}
	} else {
		fmt.Printf("\nProgram bigger than 6502 addressable RAM (64KB). Exiting.\n\n")
		os.Exit(0)
	}

	// // Load ROM to memory
	// for i := 0; i < len(data); i++ {
	// 	// F000 - F7FF (2KB Cartridge ROM)
	// 	// Memory[i] = data[i]
	// 	// F800 - FFFF (2KB Mirror Cartridge ROM)
	// 	Memory[i+0x200] = data[i]
	// }

	// // Print Memory -  Fist 2kb
	// for i := 0xF7F0; i <= 0xF7FF; i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// fmt.Println()
	// //
	// for i := 0xFFF0; i <= 0xFFFF; i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// fmt.Println()

	// // Print Memory
	// for i := 0; i < len(VGS.Memory); i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// os.Exit(2)
}

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

// Decode opcode for debug messages
func Debug_decode_console(bytes byte, mem_addr uint16) (string, string, string) {

	var (
		opcode_string            string
		operand_string           string
		operand_bigendian_string string
	)

	// Operator (opcode)
	opcode_string = fmt.Sprintf("%02x", Memory[mem_addr])

	// Decode operators
	for i := 1; i < int(bytes); i++ {
		operand_string += fmt.Sprintf("%02x", Memory[mem_addr+uint16(i)])
	}

	// Decode operators (big endian)
	for i := int(bytes) - 1; i >= 1; i-- {
		operand_bigendian_string += fmt.Sprintf("%02x", Memory[mem_addr+uint16(i)])
	}

	return opcode_string, operand_string, operand_bigendian_string
}

// Print internal opcode cycle in debug mode
func debugInternalOpcCycle(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, Opc_cycle_count, opc_cycles)
	}
}

// Print internal opcode cycle in debug mode - instructions with extra cycle
func debugInternalOpcCycleExtras(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", Cycle, Opc_cycle_count, opc_cycles+Opc_cycle_extra, opc_cycles, Opc_cycle_extra)
	}
}

// Print internal opcode cycle in debug mode - Branches
func debugInternalOpcCycleBranch(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + 1 cycle for branch + %d extra cycles for branch in different page)\n", Cycle, Opc_cycle_count, opc_cycles+Opc_cycle_extra+1, opc_cycles, Opc_cycle_extra)
	}
}

// Reset the internal opcode cycle counters
func resetIntOpcCycleCounters() {
	// Reset Opcode Cycle counter
	Opc_cycle_count = 1

	// Reset Opcode Extra Cycle counter
	Opc_cycle_extra = 0

	NewInstruction = true

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
