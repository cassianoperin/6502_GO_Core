package CLI

import (
	"6502/CONSOLE"
	"6502/CORE"
	"flag"
	"fmt"
	"os"
)

func CheckArgs() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [options] ROM_FILE\n\n%s -help for a list of available options\n\n", os.Args[0], os.Args[0])
		os.Exit(0)
	}

	cliHelp := flag.Bool("help", false, "Show this menu")
	cliConsole := flag.Bool("console", false, "Open program in interactive console")
	cliDebug := flag.Bool("debug", false, "Enable Debug Mode")
	cliPause := flag.Bool("pause", false, "Start emulation Paused")
	cliPC := flag.String("register_PC", "", "Set the Program Counter Address (hexadecimal)")
	flag.Parse()

	// Fisrt ensure that there is an last argument (rom name)
	if len(flag.Args()) != 0 {
		// Check if file exist
		testFile(flag.Arg(0))
	} else {
		fmt.Printf("Usage: %s [options] ROM_FILE\n  -console\n    	Open program in interactive console\n  -debug\n    	Enable Debug Mode\n  -help\n    	Show this menu\n  -pause\n    	Start emulation Paused\n  -register_PC\n    	Set the Program Counter Address (Hexadecimal)\n\n", os.Args[0])
		os.Exit(0)
	}

	// After, check the flags
	if *cliHelp {
		fmt.Printf("Usage: %s [options] ROM_FILE\n  -console\n    	Open program in interactive console\n  -debug\n    	Enable Debug Mode\n  -help\n    	Show this menu\n  -pause\n    	Start emulation Paused\n  -register_PC\n    	Set the Program Counter Address (Hexadecimal)\n\n", os.Args[0])
		os.Exit(0)
	}

	// Debug
	if *cliDebug {
		CORE.Debug = true
	}

	// PC
	if *cliPC != "" {

		// Check if input is Decimar of Hexadecimal and convert to integer
		output, error_flag := CONSOLE.Console_Hex_or_Dec(*cliPC)

		if error_flag {
			fmt.Printf("Invalid CLI \"register_PC\" value. Exiting.\n\n")
			os.Exit(0)
		} else {
			CORE.PC_as_argument = uint16(output)
		}

	}

	// Pause
	if *cliPause {
		CORE.Pause = true
	}

	// Console Mode
	if *cliConsole {

		if *cliPause {
			fmt.Printf("Console mode doesn't support Pause option.\n")
			os.Exit(0)
		}

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

		CONSOLE.StartConsole()

		// Exit program after close console
		os.Exit(0)
	}

}

func testFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File '%s' not found.\n\n", flag.Arg(0))
		os.Exit(0)
	}
}
