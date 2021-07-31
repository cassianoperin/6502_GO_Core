package CONSOLE

import (
	"6502/CORE"
	"fmt"
	"io"
	"strings"

	"github.com/cassianoperin/pseudo-terminal-go/terminal"
)

type instructuction struct {
	code        byte
	bytes       byte
	cycles      byte
	description string
	memory_mode string
}

type breakpoint struct {
	location string
	value    uint64
}

var (
	step_limit       int    = 1000
	step_debug_start uint64 = 0
	run_limit        int    = 1000
	goto_limit       int    = 1000
	breakpoints      []breakpoint
	opcode_map       = []instructuction{
		{0x0A, 1, 2, "ASL", "accumulator"},
		{0x18, 1, 2, "CLC", "implied"},
		{0xD8, 1, 2, "CLD", "implied"},
		{0x58, 1, 2, "CLI", "implied"},
		{0xB8, 1, 2, "CLV", "implied"},
		{0xCA, 1, 2, "DEX", "implied"},
		{0x88, 1, 2, "DEY", "implied"},
		{0xE8, 1, 2, "INX", "implied"},
		{0xC8, 1, 2, "INY", "implied"},
		{0x4A, 1, 2, "LSR", "accumulator"},
		{0xEA, 1, 2, "NOP", "implied"},
		{0x2A, 1, 2, "ROL", "accumulator"},
		{0x38, 1, 2, "SEC", "implied"},
		{0xF8, 1, 2, "SED", "implied"},
		{0x78, 1, 2, "SEI", "implied"},
		{0xAA, 1, 2, "TAX", "implied"},
		{0xA8, 1, 2, "TAY", "implied"},
		{0xBA, 1, 2, "TSX", "implied"},
		{0x8A, 1, 2, "TXA", "implied"},
		{0x9A, 1, 2, "TXS", "implied"},
		{0x98, 1, 2, "TYA", "implied"},
		{0x69, 2, 2, "ADC", "immediate"},
		{0x65, 2, 3, "ADC", "zeropage"},
		{0x75, 2, 4, "ADC", "zeropage,X"},
		{0x6D, 3, 4, "ADC", "absolute"},
		{0x7D, 3, 4, "ADC", "absolute,X"},
		{0x79, 3, 4, "ADC", "absolute,Y"},
		{0x61, 2, 6, "ADC", "(indirect,X)"},
		{0x71, 2, 5, "ADC", "(indirect),Y"},
		{0x29, 2, 2, "AND", "immediate"},
		{0x25, 2, 3, "AND", "zeropage"},
		{0x35, 2, 4, "AND", "zeropage,X"},
		{0x2D, 3, 4, "AND", "absolute"},
		{0x3D, 3, 4, "AND", "absolute,X"},
		{0x39, 3, 4, "AND", "absolute,Y"},
		{0x21, 2, 6, "AND", "(indirect,X)"},
		{0x31, 2, 5, "AND", "(indirect),Y"},
		{0x2C, 3, 4, "BIT", "absolute"},
		{0x24, 2, 3, "BIT", "zeropage"},
		{0xC5, 2, 3, "CMP", "zeropage"},
		{0xC9, 2, 2, "CMP", "immediate"},
		{0xD5, 2, 4, "CMP", "zeropage,X"},
		{0xCD, 3, 4, "CMP", "absolute"},
		{0xD9, 3, 4, "CMP", "absolute,Y"},
		{0xDD, 3, 4, "CMP", "absolute,X"},
		{0xD1, 2, 5, "CMP", "(indirect),Y"},
		{0xC1, 2, 6, "CMP", "(indirect,X)"},
		{0xE0, 2, 2, "CPX", "immediate"},
		{0xE4, 2, 3, "CPX", "zeropage"},
		{0xEC, 3, 4, "CPX", "absolute"},
		{0xC0, 2, 2, "CPY", "immediate"},
		{0xC4, 2, 3, "CPY", "zeropage"},
		{0xCC, 3, 4, "CPY", "absolute"},
		{0x49, 2, 2, "EOR", "immediate"},
		{0x45, 2, 3, "EOR", "zeropage"},
		{0x55, 2, 4, "EOR", "zeropage,X"},
		{0x4D, 3, 4, "EOR", "absolute"},
		{0x5D, 3, 4, "EOR", "absolute,X"},
		{0x59, 3, 4, "EOR", "absolute,Y"},
		{0x41, 2, 6, "EOR", "(indirect,X)"},
		{0x51, 2, 5, "EOR", "(indirect),Y"},
		{0xA9, 2, 2, "LDA", "immediate"},
		{0xA5, 2, 3, "LDA", "zeropage"},
		{0xB9, 3, 4, "LDA", "absolute,Y"},
		{0xBD, 3, 4, "LDA", "absolute,X"},
		{0xB1, 2, 5, "LDA", "(indirect),Y"},
		{0xB5, 2, 4, "LDA", "zeropage,X"},
		{0xAD, 3, 4, "LDA", "absolute"},
		{0xA1, 2, 6, "LDA", "(indirect,X)"},
		{0xA2, 2, 2, "LDX", "immediate"},
		{0xA6, 2, 3, "LDX", "zeropage"},
		{0xB6, 2, 4, "LDX", "zeropage,Y"},
		{0xBE, 3, 4, "LDX", "absolute,Y"},
		{0xAE, 3, 4, "LDX", "absolute"},
		{0xA0, 2, 2, "LDY", "immediate"},
		{0xA4, 2, 3, "LDY", "zeropage"},
		{0xB4, 2, 4, "LDY", "zeropage,X"},
		{0xAC, 3, 4, "LDY", "absolute"},
		{0xBC, 3, 4, "LDY", "absolute,X"},
		{0x09, 2, 2, "ORA", "immediate"},
		{0x05, 2, 3, "ORA", "zeropage"},
		{0x15, 2, 4, "ORA", "zeropage,X"},
		{0x0D, 3, 4, "ORA", "absolute"},
		{0x1D, 3, 4, "ORA", "absolute,X"},
		{0x19, 3, 4, "ORA", "absolute,Y"},
		{0x01, 2, 6, "ORA", "(indirect,X)"},
		{0x11, 2, 5, "ORA", "(indirect),Y"},
		{0xE9, 2, 2, "SBC", "immediate"},
		{0xE5, 2, 3, "SBC", "zeropage"},
		{0xF5, 2, 4, "SBC", "zeropage,X"},
		{0xED, 3, 4, "SBC", "absolute"},
		{0xFD, 3, 4, "SBC", "absolute,X"},
		{0xF9, 3, 4, "SBC", "absolute,Y"},
		{0xE1, 2, 6, "SBC", "(indirect,X)"},
		{0xF1, 2, 5, "SBC", "(indirect),Y"},
		{0x95, 2, 4, "STA", "zeropage,X"},
		{0x85, 2, 3, "STA", "zeropage"},
		{0x99, 3, 5, "STA", "absolute,Y"},
		{0x8D, 3, 4, "STA", "absolute"},
		{0x91, 2, 6, "STA", "(indirect),Y"},
		{0x9D, 3, 5, "STA", "absolute,X"},
		{0x81, 2, 6, "STA", "(indirect,X)"},
		{0x86, 2, 3, "STX", "zeropage"},
		{0x96, 2, 4, "STX", "zeropage,Y"},
		{0x8E, 3, 4, "STX", "absolute"},
		{0x84, 2, 3, "STY", "zeropage"},
		{0x94, 2, 4, "STY", "zeropage,X"},
		{0x8C, 3, 4, "STY", "absolute"},
		{0x06, 2, 5, "ASL", "zeropage"},
		{0x16, 2, 6, "ASL", "zeropage,X"},
		{0x0E, 3, 6, "ASL", "absolute"},
		{0x1E, 3, 7, "ASL", "absolute,X"},
		{0xC6, 2, 5, "DEC", "zeropage"},
		{0xD6, 2, 6, "DEC", "zeropage,X"},
		{0xCE, 3, 6, "DEC", "absolute"},
		{0xDE, 3, 7, "DEC", "absolute,X"},
		{0xE6, 2, 5, "INC", "zeropage"},
		{0xF6, 2, 6, "INC", "zeropage,X"},
		{0xEE, 3, 6, "INC", "absolute"},
		{0xFE, 3, 7, "INC", "absolute,X"},
		{0x46, 2, 5, "LSR", "zeropage"},
		{0x56, 2, 6, "LSR", "zeropage,X"},
		{0x4E, 3, 6, "LSR", "absolute"},
		{0x5E, 3, 7, "LSR", "absolute,X"},
		{0x26, 2, 5, "ROL", "zeropage"},
		{0x36, 2, 6, "ROL", "zeropage,X"},
		{0x2E, 3, 6, "ROL", "absolute"},
		{0x3E, 3, 7, "ROL", "absolute,X"},
		{0x6A, 1, 2, "ROR", "accumulator"},
		{0x66, 2, 5, "ROR", "zeropage"},
		{0x76, 2, 6, "ROR", "zeropage,X"},
		{0x6E, 3, 6, "ROR", "absolute"},
		{0x7E, 3, 7, "ROR", "absolute,X"},
		{0x48, 1, 3, "PHA", "implied"},
		{0x08, 1, 3, "PHP", "implied"},
		{0x68, 1, 4, "PLA", "implied"},
		{0x28, 1, 4, "PLP", "implied"},
		{0x4C, 3, 3, "JMP", "absolute"},
		{0x6C, 3, 5, "JMP", "indirect"},
		{0x20, 3, 6, "JSR", "absolute"},
		{0x40, 1, 6, "RTI", "implied"},
		{0x60, 1, 6, "RTS", "implied"},
		{0x00, 1, 7, "BRK", "implied"},
		{0xD0, 2, 2, "BNE", "relative"},
		{0xF0, 2, 2, "BEQ", "relative"},
		{0x10, 2, 2, "BPL", "relative"},
		{0x30, 2, 2, "BMI", "relative"},
		{0x70, 2, 2, "BVS", "relative"},
		{0x50, 2, 2, "BVC", "relative"},
		{0xB0, 2, 2, "BCS", "relative"},
		{0x90, 2, 2, "BCC", "relative"},
	}
)

func StartConsole() {

	// Reset system
	CORE.Reset()

	term, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}
	defer term.ReleaseFromStdInOut() // defer this
	fmt.Printf("\n------------------------ Console mode ------------------------\n\nType \"Ctrl-Q\" to quit, \"help\" for more options\n\n")

	// Print Header
	Console_PrintHeader()

	term.SetPrompt("# ")
	line, err := term.ReadLine()
	for {
		if err == io.EOF {
			term.Write([]byte(line))
			fmt.Println()
			return
		}
		if (err != nil && strings.Contains(err.Error(), "control-c break")) || len(line) == 0 {
			line, err = term.ReadLine()
		} else {

			// Command Interpreter
			CommandInterpreter(line)

			// term.Write([]byte(line + "\r\n"))
			line, err = term.ReadLine()
		}
	}
	// term.Write([]byte(line))
}

// Interpreter
func CommandInterpreter(text string) {

	// Remove duplicated spaces
	text = strings.Join(strings.Fields(strings.TrimSpace(text)), " ")

	// Convert to Lower Case
	text = strings.ToLower(text)

	// Split commands and arguments by Spaces
	text_slice := strings.Split(text, " ")

	if text_slice[0] == "quit" || text_slice[0] == "exit" { // QUIT

		Console_Command_Quit()

	} else if text_slice[0] == "help" { // HELP
		Console_Command_Help()

	} else if text_slice[0] == "step" { // STEP

		Console_Command_Step(text_slice)

	} else if text_slice[0] == "step_limit" { // STEP_LIMIT

		Console_Command_StepLimit(text_slice)

	} else if text_slice[0] == "step_debug_start" { // STEP_DEBUG_START

		Console_Command_StepDebugStart(text_slice)

	} else if text_slice[0] == "add_breakpoint" { // ADD BREAKPOINT

		Console_Command_AddBreakpoint(text_slice)

	} else if text_slice[0] == "del_breakpoint" { // DELETE BREAKPOINT

		Console_Command_DelBreakpoint(text_slice)

	} else if text_slice[0] == "show_breakpoints" { // SHOW BREAKPOINTS

		Console_Command_ShowBreakpoint(text_slice)

	} else if text_slice[0] == "run" { // RUN

		Console_Command_Run(text_slice)

	} else if text_slice[0] == "run_limit" { // RUN LIMIT

		Console_Command_RunLimit(text_slice)

	} else if text_slice[0] == "mem" { // MEMORY

		Console_Command_Mem(text_slice)

	} else if text_slice[0] == "goto" { // GOTO

		Console_Command_Goto(text_slice)

	} else if text_slice[0] == "goto_limit" { // GOTO_LIMIT

		Console_Command_GotoLimit(text_slice)

	} else if text_slice[0] == "reset" { // RESET

		Console_Command_Reset(text_slice)

	} else if text_slice[0] == "disassemble" { // DISASSEMBLE

		Console_Command_Disassemble(text_slice)

	} else if text_slice[0] == "registers" { // REGISTERS

		Console_Command_Registers(text_slice)

	} else if text_slice[0] == "processor_status" { // PROCESSOR STATUS

		Console_Command_ProcessorStatus(text_slice)

	} else if text_slice[0] == "debug" { // DEBUG

		Console_Command_Debug(text_slice)

	} else { // Command not found
		fmt.Printf("Command not found\n\n")
	}

}
