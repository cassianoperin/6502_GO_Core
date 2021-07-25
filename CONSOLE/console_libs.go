package CONSOLE

import (
	"6502/CORE"
	"fmt"
	"strconv"
	"strings"
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
				if CORE.PC == uint16(breakpoints[i].value) {
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

func Console_Hex_or_Dec(value string) (int, bool) {

	var (
		mem1       int
		error_flag bool
	)

	// Test if the value start if 0x or 0X
	if strings.HasPrefix(value, "0x") || strings.HasPrefix(value, "0X") {

		// HEXADECIMAL Input

		var hexaString string = value
		numberStr := strings.Replace(hexaString, "0x", "", 1)
		numberStr = strings.Replace(numberStr, "0X", "", 1)

		tmp_value, err := strconv.ParseInt(numberStr, 16, 64)

		if err != nil {
			// fmt.Println("Invalid value.")
			error_flag = true
		} else {
			// Convert to decimal and set mem1 value
			mem1 = int(tmp_value)
		}

	} else {

		// DECIMAL Input

		tmp_value, err := strconv.Atoi(value)
		if err != nil {
			// handle error
			// fmt.Printf("Invalid value %s\n\n", value)
			error_flag = true

		} else {
			// Set mem1 value
			mem1 = int(tmp_value)
		}
	}

	return mem1, error_flag

}
