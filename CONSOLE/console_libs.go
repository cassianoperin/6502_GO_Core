package CONSOLE

import (
	"6502/CORE"
	"fmt"
)

// --------------------------------------- Debug ---------------------------------------- //

// Print the debub information in console
func print_debug_console(opcode_map []instructuction) {
	for i := 0; i < len(opcode_map); i++ {

		if CORE.Memory[CORE.PC] == opcode_map[i].code {
			opc_string, opc_operand := CORE.Debug_decode_console(opcode_map[i].bytes)
			fmt.Printf("\n--> $%04X\t%s %s\t%s %s (%s)\n\n", CORE.PC, opc_string, opc_operand, opcode_map[i].description, opc_operand, opcode_map[i].memory_mode)
		}

	}
}

// Print Help Menu
func Console_PrintHelp() {
	// fmt.Printf("\n\tquit\t\t\t\t\t\tQuit console\n\thelp\t\t\t\t\t\tPrint help menu\n\t--\n\tstep\t\t\t\t\t\tExecute current opcode\n\tstep <value>\t\t\t\t\tExecute <value> opcodes\n\tstep_limit <value>\t\t\t\tDefine the maximum steps allowed\n\t--\n\tadd_breakpoint <PC|A|X|Y|CYCLE>=<Value>\t\tAdd a breakpoint\n\tdel_breakpoint <index>\t\t\t\tDelete a breakpoint\n\tshow_breakpoints\t\t\t\tShow breakpoints\n\t--\n\tmem || mem <address> || mem <address> <address>\tDump address values\n\t--\n\trun\t\t\t\t\t\tRun the emulator\n\trun_limit <value>\t\t\t\tDefine the maximum steps allowed in RUN\n\n")
	fmt.Printf("\n\tquit\t\t\t\t\t\tQuit console\n\thelp\t\t\t\t\t\tPrint help menu\n\t-")
	fmt.Printf("\n\tstep\t\t\t\t\t\tExecute current opcode\n\tstep <value>\t\t\t\t\tExecute <value> opcodes\n\tstep_limit <value>\t\t\t\tDefine the maximum steps allowed\n\t-")
	fmt.Printf("\n\tadd_breakpoint <PC|A|X|Y|CYCLE>=<Value>\t\tAdd a breakpoint\n\tdel_breakpoint <index>\t\t\t\tDelete a breakpoint\n\tshow_breakpoints\t\t\t\tShow breakpoints\n\t-")
	fmt.Printf("\n\tmem\t\t\t\t\t\tDump full memory\n\tmem <address>\t\t\t\t\tDump memory address\n\tmem <start address> <end address>\t\tDump memory address range\n\t-")
	fmt.Printf("\n\trun\t\t\t\t\t\tRun the emulator\n\trun_limit <value>\t\t\t\tDefine the maximum steps allowed in RUN\n\n")
}

func Console_PrintAddBrkErr() {
	fmt.Printf("Usage: add_breakpoint < PC | A | X | Y | CYCLE > = <Value>\n\n")
}

func Console_PrintStepLimitErr() {
	fmt.Printf("Usage: step_limit <Value>\n\n")
}

func Console_PrintRunLimitErr() {
	fmt.Printf("Usage: run_limit <Value>\n\n")
}

// Print current Processor Information
func Console_PrintHeader() {
	fmt.Printf("  -------------------------------------------------------------------------\n")
	fmt.Printf("  |   PC\tA\tX\tY\tSP\tNV-BDIZC\tCycle\n")
	fmt.Printf("  |   %04X\t%02X\t%02X\t%02X\t%02X\t%d%d%d%d%d%d%d%d\t%d\n", CORE.PC, CORE.A, CORE.X, CORE.Y, CORE.SP, CORE.P[7], CORE.P[6], CORE.P[5], CORE.P[4], CORE.P[3], CORE.P[2], CORE.P[1], CORE.P[0], CORE.Cycle)
	fmt.Printf("  -------------------------------------------------------------------------\n\n")
}

// ---------------------------------------- Libs ---------------------------------------- //

// Execute the necessary cycles for next instruction
func Console_Step(opcode_map []instructuction) {
	// Print the opcode debug
	print_debug_console(opcode_map)

	for !CORE.NewInstruction {
		CORE.CPU_Interpreter()
	}

	// Reset new instruction flag
	CORE.NewInstruction = false

	// Print the Header
	Console_PrintHeader()
}

// Execute the necessary cycles for next instruction without print on console
func Console_Step_without_debug(opcode_map []instructuction) {

	for !CORE.NewInstruction {
		CORE.CPU_Interpreter()
	}

	// Reset new instruction flag
	CORE.NewInstruction = false
}

// Remove a value from a slice
func Console_Remove_breakpoint(s []breakpoint, index int) []breakpoint {
	return append(s[:index], s[index+1:]...)
}

func Console_Check_breakpoints(break_flag bool) bool {
	// Check Breakpoints
	if len(breakpoints) > 0 {
		for i := 0; i < len(breakpoints); i++ {

			// ------ PC ------ //
			if breakpoints[i].location == "PC" {
				if CORE.PC == breakpoints[i].value {
					fmt.Printf("Breakpoint %d reached: %s=0x%04X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ A ------- //
			} else if breakpoints[i].location == "A" {
				if CORE.A == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=0x%02X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ X ------- //
			} else if breakpoints[i].location == "X" {
				if CORE.X == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=0x%02X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ Y ------- //
			} else if breakpoints[i].location == "Y" {
				if CORE.Y == byte(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=0x%02X\t(Decimal: %d)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}

				// ------ Cycle ------- //
			} else if breakpoints[i].location == "CYCLE" {
				if CORE.Cycle >= uint64(breakpoints[i].value) {
					fmt.Printf("Breakpoint %d reached: %s=%d\t(0x%02X)\n", i, breakpoints[i].location, breakpoints[i].value, breakpoints[i].value)
					break_flag = true
				}
			}

		}

	}

	return break_flag
}
