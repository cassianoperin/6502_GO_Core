package main

import (
	"6502/CLI"
	"6502/CORE"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/faiface/pixel/pixelgl"
)

var PC_arg uint16 = 0

// Function used by readROM to avoid 'bytesread' return
func ReadContent(file *os.File, bytes_number int) []byte {

	bytes := make([]byte, bytes_number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

// Read ROM and write it to the RAM
func readROM(filename string) {

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

	if romsize == 65536 {
		// Load ROM to memory
		for i := 0; i < len(data); i++ {
			// F000 - F7FF (2KB Cartridge ROM)
			CORE.Memory[i] = data[i]
			// F800 - FFFF (2KB Mirror Cartridge ROM)
			CORE.Memory[i] = data[i]
		}
	}

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

func main() {

	fmt.Printf("\nMOS 6502 CPU Emulator\n\n")

	// Validate the Arguments
	CLI.CheckArgs()

	// Set initial variables values
	CORE.Initialize()
	// Initialize Timers
	CORE.InitializeTimers()

	// Read ROM to the memory
	// readROM(os.Args[1])
	readROM(flag.Arg(0))
	// readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_functional_test.bin")
	// readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_decimal_test.bin")

	// Reset system
	CORE.Reset()

	// Overwrite PC if requested in arguments
	if PC_arg != 0 {
		CORE.PC = PC_arg
	}

	// Start Window System and draw Graphics
	pixelgl.Run(CORE.Run)

}
