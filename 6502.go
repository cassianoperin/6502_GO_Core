package main

import (
	"6502/CLI"
	"6502/CORE"
	"flag"
	"fmt"

	"github.com/faiface/pixel/pixelgl"
)

func main() {

	fmt.Printf("\nMOS 6502 CPU Emulator\n\n")

	// Validate the Arguments
	CLI.CheckArgs()

	// Set initial variables values
	CORE.Initialize()

	// Initialize Timers
	CORE.InitializeTimers()

	// Read ROM to the memory
	CORE.ReadROM(flag.Arg(0))
	// readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_functional_test.bin")
	// readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_decimal_test.bin")

	// Reset system
	CORE.Reset()

	// Start Window System and draw Graphics
	pixelgl.Run(CORE.Run)

}
