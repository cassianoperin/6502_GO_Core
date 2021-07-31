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
	cliDebug := flag.Bool("debug", false, "Enable Debug Mode")
	cliPC := flag.String("register_PC", "", "Set the Program Counter Address (hexadecimal)")
	flag.Parse()

	// Fisrt ensure that there is an last argument (rom name)
	if len(flag.Args()) != 0 {
		// Check if file exist
		testFile(flag.Arg(0))
	} else {
		fmt.Printf("Usage: %s [options] ROM_FILE\n  -debug\n    	Enable Debug Mode\n  -help\n    	Show this menu\n  -register_PC\n    	Set the Program Counter Address (Hexadecimal)\n\n", os.Args[0])
		os.Exit(0)
	}

	// After, check the flags
	if *cliHelp {
		fmt.Printf("Usage: %s [options] ROM_FILE\n  -debug\n    	Enable Debug Mode\n  -help\n    	Show this menu\n  -register_PC\n    	Set the Program Counter Address (Hexadecimal)\n\n", os.Args[0])
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

}

func testFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File '%s' not found.\n\n", flag.Arg(0))
		os.Exit(0)
	}
}
