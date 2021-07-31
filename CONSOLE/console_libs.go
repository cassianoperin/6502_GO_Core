package CONSOLE

import (
	"6502/CORE"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --------------------------------------- Debug ---------------------------------------- //

// Print the debub information in console for step and disassemble operation
func print_debug_console(opcode_map []instructuction, mem_arg int) {

	var mem_string string
	var opcode_found bool

	for i := 0; i < len(opcode_map); i++ {

		if CORE.Memory[mem_arg] == opcode_map[i].code {

			opcode_found = true

			opc_string, opc_operand, operand_bigendian_string := CORE.Debug_decode_console(opcode_map[i].bytes, uint16(mem_arg))

			// Map Opcode
			switch opcode_map[i].memory_mode {

			case "implied":
				mem_string = "    "

			case "accumulator":
				mem_string = "A   "

			case "relative", "":
				mem_string = "$" + operand_bigendian_string + "  "

			case "immediate":
				mem_string = "#$" + operand_bigendian_string

			case "absolute":
				mem_string = "$" + operand_bigendian_string

			case "absolute,X":
				mem_string = "$" + operand_bigendian_string + ",X"

			case "absolute,Y":
				mem_string = "$" + operand_bigendian_string + ",Y"

			case "zeropage":
				mem_string = "$" + operand_bigendian_string + "  "

			case "zeropage,X":
				mem_string = "$" + operand_bigendian_string + ",X"

			case "zeropage,Y":
				mem_string = "$" + operand_bigendian_string + ",Y"

			case "indirect":
				mem_string = "($" + operand_bigendian_string + ")"

			case "(indirect,X)":
				mem_string = "($" + operand_bigendian_string + ",X)"

			case "(indirect),Y":
				mem_string = "($" + operand_bigendian_string + "),Y"

			default:
				fmt.Printf("print_debug_console_disassembler(): Memory mode not mapped\n\n")
				os.Exit(2)
			}

			fmt.Printf("\t$%04X\t%s %s\t\t%s %s\t(%s)\n", mem_arg, opc_string, opc_operand, opcode_map[i].description, mem_string, opcode_map[i].memory_mode)

		}

	}

	if !opcode_found {
		fmt.Printf("\t$%04X\t???\n", mem_arg)
	}
}

func Console_PrintAddBrkErr() {
	fmt.Printf("Usage: add_breakpoint < PC | A | X | Y | CYCLE > = <Value>\n\n")
}

func Console_PrintStepLimitErr() {
	fmt.Printf("Usage: step_limit <Value>\n\n")
}

func Console_PrintGotoLimitErr() {
	fmt.Printf("Usage: goto_limit <Value>\n\n")
}

func Console_PrintRunLimitErr() {
	fmt.Printf("Usage: run_limit <Value>\n\n")
}

// Print current Processor Information
func Console_PrintHeader() {
	fmt.Printf("\n\n\n")
	fmt.Printf("  -------------------------------------------------------------------------\n")
	fmt.Printf("  |   PC\tA\tX\tY\tSP\tNV-BDIZC\tCycle\n")
	fmt.Printf("  |   %04X\t%02X\t%02X\t%02X\t%02X\t%d%d%d%d%d%d%d%d\t%d\n", CORE.PC, CORE.A, CORE.X, CORE.Y, CORE.SP, CORE.P[7], CORE.P[6], CORE.P[5], CORE.P[4], CORE.P[3], CORE.P[2], CORE.P[1], CORE.P[0], CORE.Cycle)
	fmt.Printf("  -------------------------------------------------------------------------\n\n")
}

// ---------------------------------------- Libs ---------------------------------------- //

// Execute the necessary cycles for next instruction
func Console_Step(opcode_map []instructuction, origin_command string) {

	// Keep current debug value
	current_debug := CORE.Debug

	// Debug is only used in STEP command
	if origin_command == "run" || origin_command == "goto" {
		CORE.Debug = false // Force disable debug
	}

	// Print the opcode debug
	fmt.Println()
	print_debug_console(opcode_map, int(CORE.PC))
	fmt.Println()

	for !CORE.NewInstruction {
		CORE.CPU_Interpreter()
	}

	// Reset new instruction flag
	CORE.NewInstruction = false

	// Print the Header
	Console_PrintHeader()

	// Debug is only used in STEP command
	if origin_command == "run" || origin_command == "goto" {
		CORE.Debug = current_debug // // Return original Debug value
	}
}

// Execute the necessary cycles for next instruction without print on console
func Console_Step_without_debug(opcode_map []instructuction, origin_command string) {

	// Keep current debug value
	current_debug := CORE.Debug

	// Debug is only used in STEP command
	if origin_command == "run" || origin_command == "goto" {
		CORE.Debug = false // Force disable debug
	}

	for !CORE.NewInstruction {
		CORE.CPU_Interpreter()
	}

	// Reset new instruction flag
	CORE.NewInstruction = false

	// Debug is only used in STEP command
	if origin_command == "run" || origin_command == "goto" {
		CORE.Debug = current_debug // // Return original Debug value
	}

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
