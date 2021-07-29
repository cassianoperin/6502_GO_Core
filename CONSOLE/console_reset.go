package CONSOLE

import (
	"6502/CORE"
	"flag"
	"fmt"
)

// Print Help Menu
func Console_Command_Reset(text_slice []string) {
	// Test the command syntax
	if len(text_slice) > 1 {

		// Show current value
		fmt.Printf("Reset doesn't accept arguments\n\n")

	} else {

		// Reset CPU
		CORE.Initialize()
		CORE.Reset()

		// Load Program again into memory
		CORE.ReadROM(flag.Arg(0))

		// Print the Header
		Console_PrintHeader()
	}
}
