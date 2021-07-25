package CONSOLE

import (
	"fmt"
	"os"
)

// Quit Console
func Console_Command_Quit() {
	fmt.Printf("Exiting console.\n")
	os.Exit(0)
}
