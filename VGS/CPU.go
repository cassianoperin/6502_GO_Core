package VGS

import (
	"fmt"
	"os"
	"time"
)

// Initialization
func Initialize() {

	// Clean Memory Array
	Memory = [65536]byte{}
	// Clean CPU Variables
	PC = 0
	opcode = 0
	X = 0
	Y = 0
	A = 0
	P = [8]byte{}

	// Initialize CPU
	CPU_Enabled = true

	// Initialize P (Bit 4 (Break) and Bit 5 (Unused))
	P[5] = 1 // Always set

	// Atari 2600 interpreter mode
	if CPU_MODE == 0 {

		// Always Enabled since 6507 doesn't have interrupts
		P[4] = 1

		// 6502/6507 interpreter mode
	} else {

		// Will be set with BRK instruction
		P[4] = 0

	}

}

func InitializeTimers() {
	// Start Timers
	clock_timer = time.NewTicker(time.Nanosecond)          // CPU Clock
	screenRefresh_timer = time.NewTicker(time.Second / 30) // 60Hz Clock for screen refresh rate
}

// Reset Vector // 0xFFFC | 0xFFFD (Little Endian)
func Reset() {

	// Atari 2600 interpreter mode
	if CPU_MODE == 0 {

		// Read the Opcode from PC+1 and PC bytes (Little Endian)
		PC = uint16(Memory[0xFFFD])<<8 | uint16(Memory[0xFFFC])

		// 6502/6507 interpreter mode
	} else {

		// Set the PC on the start of programs
		PC = 0x400
	}

	// Reset the SP
	SP = 0xFF
}

func Show() {
	if Debug {
		fmt.Printf("\n\n%04X : %02X\n\n", PC, opcode)
		fmt.Printf("\t\t\t\t\t\t\t\t\t\t   N V - B D I Z C")
		fmt.Printf("\nCycle: %d\tOpcode: %02X\tPC: 0x%04X(%d)\tA: 0x%02X\tX: 0x%02X\tY: 0x%02X\tP: %d %d %d %d %d %d %d %d\tSP: %02X\t\tStack:  Mem[1FF]: %02X   Mem[1FE]: %02X   Mem[1FD]: %02X   Mem[1FC]: %02X\n", counter_F_Cycle, opcode, PC, PC, A, X, Y, P[7], P[6], P[5], P[4], P[3], P[2], P[1], P[0], SP, Memory[0x1FF], Memory[0x1FE], Memory[0x1FD], Memory[0x1FC])
	}
}

// CPU Interpreter
func CPU_Interpreter() {

	// Read the Next Instruction to be executed
	opcode = Memory[PC]

	// Print Cycle and Debug Information
	// if Debug {
	// Just show in the first opcode cycle
	if opc_cycle_count == 1 {
		Show()
	}
	// }

	// Map Opcode
	switch opcode {

	//-------------------------------------------------- Implied --------------------------------------------------//

	case 0x78: // Instruction SEI
		opc_SEI(1, 2)

	case 0x58: // Instruction CLI
		opc_CLI(1, 2)

	case 0xB8: // Instruction CLV
		opc_CLV(1, 2)

	case 0x38: // Instruction SEC
		opc_SEC(1, 2)

	case 0xF8: // Instruction SED
		opc_SED(1, 2)

	case 0x18: // Instruction CLC
		opc_CLC(1, 2)

	case 0xD8: // Instruction CLD
		opc_CLD(1, 2)

	case 0x8A: // Instruction TXA
		opc_TXA(1, 2)

	case 0x98: // Instruction TYA
		opc_TYA(1, 2)

	case 0xAA: // Instruction TAX
		opc_TAX(1, 2)

	case 0xA8: // Instruction TAY
		opc_TAY(1, 2)

	case 0xCA: // Instruction DEX
		opc_DEX(1, 2)

	case 0x88: // Instruction DEY
		opc_DEY(1, 2)

	case 0x9A: // Instruction TXS
		opc_TXS(1, 2)

	case 0x48: // Instruction PHA
		opc_PHA(1, 3)

	case 0x28: // Instruction PLP
		opc_PLP(1, 4)

	case 0x08: // Instruction PHP
		opc_PHP(1, 3)

	case 0x68: // Instruction PLA
		opc_PLA(1, 4)

	case 0x00: // Instruction BRK
		opc_BRK(1, 7)

	case 0x40: // Instruction RTI
		opc_RTI(1, 6)

	case 0xC8: // Instruction INY
		opc_INY(1, 2)

	case 0xE8: // Instruction INX
		opc_INX(1, 2)

	case 0x60: // Instruction RTS
		opc_RTS(1, 6)

	case 0xEA: // Instruction NOP
		opc_NOP(1, 2)

	case 0xBA: // Instruction TSX
		opc_TSX(1, 2)

	//-------------------------------------------------- Just zeropage --------------------------------------------------//

	case 0xE6: // Instruction INC ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_INC(memAddr, memMode, 2, 5)

	case 0xF6: // Instruction INC ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_INC(memAddr, memMode, 2, 6)

	case 0xEE: // Instruction INC ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_INC(memAddr, memMode, 3, 6)

	case 0xFE: // Instruction INC ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_INC(memAddr, memMode, 3, 7)

	//-------------------------------------------- Branches - just relative ---------------------------------------------//

	case 0xD0: // Instruction BNE ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BNE(memValue, 2, 2)

	case 0x90: // Instruction BCC ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BCC(memValue, 2, 2)

	case 0x50: // Instruction BVC ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BVC(memValue, 2, 2)

	case 0xB0: // Instruction BCS ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BCS(memValue, 2, 2)

	case 0x30: // Instruction BMI ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BMI(memValue, 2, 2)

	case 0x10: // Instruction BPL ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BPL(memValue, 2, 2)

	case 0xF0: // Instruction BEQ ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BEQ(memValue, 2, 2)

	case 0x70: // Instruction BVS ( relative )
		if opc_cycle_count == 1 {
			memValue = addr_mode_Relative(PC + 1)
		}
		opc_BVS(memValue, 2, 2)

	//-------------------------------------------------- LDX --------------------------------------------------//

	case 0xA2: // Instruction LDX ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_LDX(memAddr, memMode, 2, 2)

	case 0xA6: // Instruction LDX ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_LDX(memAddr, memMode, 2, 3)

	case 0xB6: // Instruction LDX ( zeropage,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageY(PC + 1)
		}
		opc_LDX(memAddr, memMode, 2, 4)

	case 0xBE: // Instruction LDX ( absolute,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteY(PC + 1)
		}
		opc_LDX(memAddr, memMode, 3, 4)

	case 0xAE: // Instruction LDX ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_LDX(memAddr, memMode, 3, 4)

	//-------------------------------------------------- STX --------------------------------------------------//

	case 0x86: // Instruction STX ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_STX(memAddr, memMode, 2, 3)

	case 0x96: // Instruction STX ( zeropage,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageY(PC + 1)
		}
		opc_STX(memAddr, memMode, 2, 4)

	case 0x8E: // Instruction STX ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_STX(memAddr, memMode, 3, 4)

	//-------------------------------------------------- JMP --------------------------------------------------//

	case 0x4C: // Instruction JMP ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_JMP(memAddr, memMode, 3, 3)

	case 0x6C: // Instruction JMP (indirect)
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Indirect(PC + 1)
		}
		opc_JMP(memAddr, memMode, 3, 5)

	case 0x20: // Instruction JSR ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_JSR(memAddr, memMode, 3, 6)

	//-------------------------------------------------- BIT --------------------------------------------------//

	case 0x2C: // Instruction BIT ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_BIT(memAddr, memMode, 3, 4)

	case 0x24: // Instruction BIT ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_BIT(memAddr, memMode, 2, 3)

	//-------------------------------------------------- LDA --------------------------------------------------//

	case 0xA9: // Instruction LDA ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_LDA(memAddr, memMode, 2, 2)

	case 0xA5: // Instruction LDA ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_LDA(memAddr, memMode, 2, 3)

	case 0xB9: // Instruction LDA ( absolute,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteY(PC + 1)
		}
		opc_LDA(memAddr, memMode, 3, 4)

	case 0xBD: // Instruction LDA ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_LDA(memAddr, memMode, 3, 4)

	case 0xB1: // Instruction LDA ( (indirect),Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectY(PC + 1)
		}
		opc_LDA(memAddr, memMode, 2, 5)

	case 0xB5: // Instruction LDA ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_LDA(memAddr, memMode, 2, 4)

	case 0xAD: // Instruction LDA ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_LDA(memAddr, memMode, 3, 4)

	case 0xA1: // Instruction LDA ( (indirect,X) )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectX(PC + 1)
		}
		opc_LDA(memAddr, memMode, 2, 6)

	//-------------------------------------------------- LDY --------------------------------------------------//

	case 0xA0: // Instruction LDY ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_LDY(memAddr, memMode, 2, 2)

	case 0xA4: // Instruction LDY ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_LDY(memAddr, memMode, 2, 3)

	case 0xB4: // Instruction LDY ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_LDY(memAddr, memMode, 2, 4)

	case 0xAC: // Instruction LDY ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_LDY(memAddr, memMode, 3, 4)

	case 0xBC: // Instruction LDY ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_LDY(memAddr, memMode, 3, 4)

	//-------------------------------------------------- STY --------------------------------------------------//

	case 0x84: // Instruction STY ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_STY(memAddr, memMode, 2, 3)

	case 0x94: // Instruction STY ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_STY(memAddr, memMode, 2, 4)

	case 0x8C: // Instruction STY ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_STY(memAddr, memMode, 3, 4)

	//-------------------------------------------------- CPY --------------------------------------------------//

	case 0xC0: // Instruction CPY ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_CPY(memAddr, memMode, 2, 2)

	case 0xC4: // Instruction STCPYY ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_CPY(memAddr, memMode, 2, 3)

	case 0xCC: // Instruction CPY ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_CPY(memAddr, memMode, 3, 4)

	//-------------------------------------------------- CPX --------------------------------------------------//

	case 0xE0: // Instruction CPX ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_CPX(memAddr, memMode, 2, 2)

	case 0xE4: // Instruction CPX ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_CPX(memAddr, memMode, 2, 3)

	case 0xEC: // Instruction CPX ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_CPX(memAddr, memMode, 3, 4)

	//-------------------------------------------------- SBC --------------------------------------------------//

	case 0xE5: // Instruction STY ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_SBC(memAddr, memMode, 2, 3)

	case 0xE9: // Instruction STY ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_SBC(memAddr, memMode, 2, 2)

	//-------------------------------------------------- DEC --------------------------------------------------//

	case 0xC6: // Instruction DEC ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_DEC(memAddr, memMode, 2, 5)

	case 0xD6: // Instruction DEC ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_DEC(memAddr, memMode, 2, 6)

	case 0xCE: // Instruction DEC ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_DEC(memAddr, memMode, 3, 6)

	case 0xDE: // Instruction DEC ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_DEC(memAddr, memMode, 3, 7)

	//-------------------------------------------------- AND --------------------------------------------------//

	case 0x29: // Instruction AND ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_AND(memAddr, memMode, 2, 2)

	case 0x25: // Instruction AND ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_AND(memAddr, memMode, 2, 3)

	case 0x35: // Instruction AND ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_AND(memAddr, memMode, 2, 4)

	case 0x2D: // Instruction AND ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_AND(memAddr, memMode, 3, 4)

	case 0x3D: // Instruction AND ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_AND(memAddr, memMode, 3, 4)

	case 0x39: // Instruction AND ( absolute,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteY(PC + 1)
		}
		opc_AND(memAddr, memMode, 3, 4)

	case 0x21: // Instruction AND ( (indirect,X) )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectX(PC + 1)
		}
		opc_AND(memAddr, memMode, 2, 6)

	case 0x31: // Instruction AND ( (indirect),Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectY(PC + 1)
		}
		opc_AND(memAddr, memMode, 2, 5)

	//-------------------------------------------------- ORA --------------------------------------------------//

	case 0x09: // Instruction ORA ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_ORA(memAddr, memMode, 2, 2)

	case 0x05: // Instruction ORA ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_ORA(memAddr, memMode, 2, 3)

	case 0x15: // Instruction ORA ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_ORA(memAddr, memMode, 2, 4)

	case 0x0D: // Instruction ORA ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_ORA(memAddr, memMode, 3, 4)

	case 0x1D: // Instruction ORA ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_ORA(memAddr, memMode, 3, 4)

	case 0x19: // Instruction ORA ( absolute,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteY(PC + 1)
		}
		opc_ORA(memAddr, memMode, 3, 4)

	case 0x01: // Instruction ORA ( (indirect,X) )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectX(PC + 1)
		}
		opc_ORA(memAddr, memMode, 2, 6)

	case 0x11: // Instruction ORA ( (indirect),Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectY(PC + 1)
		}
		opc_ORA(memAddr, memMode, 2, 5)

	//-------------------------------------------------- EOR --------------------------------------------------//

	case 0x49: // Instruction EOR ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_EOR(memAddr, memMode, 2, 2)

	case 0x45: // Instruction EOR ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_EOR(memAddr, memMode, 2, 3)

	case 0x55: // Instruction EOR ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_EOR(memAddr, memMode, 2, 4)

	case 0x4D: // Instruction EOR ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_EOR(memAddr, memMode, 3, 4)

	case 0x5D: // Instruction EOR ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_EOR(memAddr, memMode, 3, 4)

	case 0x59: // Instruction EOR ( absolute,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteY(PC + 1)
		}
		opc_EOR(memAddr, memMode, 3, 4)

	case 0x41: // Instruction EOR ( (indirect,X) )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectX(PC + 1)
		}
		opc_EOR(memAddr, memMode, 2, 6)

	case 0x51: // Instruction EOR ( (indirect),Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectY(PC + 1)
		}
		opc_EOR(memAddr, memMode, 2, 5)

	//-------------------------------------------------- ASL --------------------------------------------------//

	case 0x0A: // Instruction ASL (  accumulator )
		opc_ASL_A(1, 2)

	case 0x06: // Instruction ASL ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_ASL(memAddr, memMode, 2, 5)

	case 0x16: // Instruction ASL ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_ASL(memAddr, memMode, 2, 6)

	case 0x0E: // Instruction ASL ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_ASL(memAddr, memMode, 3, 6)

	case 0x1E: // Instruction ASL ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_ASL(memAddr, memMode, 3, 7)

	//-------------------------------------------------- LSR --------------------------------------------------//

	case 0x4A: // Instruction LSR (  accumulator )
		opc_LSR_A(1, 2)

	case 0x46: // Instruction LSR ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_LSR(memAddr, memMode, 2, 5)

	case 0x56: // Instruction LSR ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_LSR(memAddr, memMode, 2, 6)

	case 0x4E: // Instruction LSR ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_LSR(memAddr, memMode, 3, 6)

	case 0x5E: // Instruction LSR ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_LSR(memAddr, memMode, 3, 7)

	//-------------------------------------------------- CMP --------------------------------------------------//

	case 0xC5: // Instruction CMP ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_CMP(memAddr, memMode, 2, 3)

	case 0xC9: // Instruction CMP ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_CMP(memAddr, memMode, 2, 2)

	case 0xD5: // Instruction CMP ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_CMP(memAddr, memMode, 2, 4)

	case 0xCD: // Instruction CMP ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_CMP(memAddr, memMode, 3, 4)

	case 0xD9: // Instruction CMP ( absolute,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteY(PC + 1)
		}
		opc_CMP(memAddr, memMode, 3, 4)

	case 0xDD: // Instruction CMP ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_CMP(memAddr, memMode, 3, 4)

	case 0xD1: // Instruction CMP ( (indirect),Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectY(PC + 1)
		}
		opc_CMP(memAddr, memMode, 2, 5)

	case 0xC1: // Instruction CMP ( (indirect,X) )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectX(PC + 1)
		}
		opc_CMP(memAddr, memMode, 2, 6)

	//-------------------------------------------------- STA --------------------------------------------------//

	case 0x95: // Instruction STA ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_STA(memAddr, memMode, 2, 4)

	case 0x85: // Instruction STA ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_STA(memAddr, memMode, 2, 3)

	case 0x99: // Instruction STA ( absolute,Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteY(PC + 1)
		}
		opc_STA(memAddr, memMode, 3, 5)

	case 0x8D: // Instruction STA ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_STA(memAddr, memMode, 3, 4)

	case 0x91: // Instruction STA ( (indirect),Y )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectY(PC + 1)
		}
		opc_STA(memAddr, memMode, 2, 6)

	case 0x9D: // Instruction STA ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_STA(memAddr, memMode, 3, 5)

	case 0x81: // Instruction STA ( (indirect,X) )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_IndirectX(PC + 1)
		}
		opc_STA(memAddr, memMode, 2, 6)

	//-------------------------------------------------- ADC --------------------------------------------------//

	case 0x65: // Instruction ADC ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_ADC(memAddr, memMode, 2, 3)

	case 0x7D: // Instruction ADC ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_ADC(memAddr, memMode, 3, 4)

	case 0x69: // Instruction ADC ( immediate )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Immediate(PC + 1)
		}
		opc_ADC(memAddr, memMode, 2, 2)

	//-------------------------------------------------- ROL --------------------------------------------------//

	case 0x2A: // Instruction ROL (  accumulator )
		opc_ROL_A(1, 2)

	case 0x26: // Instruction ROL ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_ROL(memAddr, memMode, 2, 5)

	case 0x36: // Instruction ROL ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_ROL(memAddr, memMode, 2, 6)

	case 0x2E: // Instruction ROL ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_ROL(memAddr, memMode, 3, 6)

	case 0x3E: // Instruction ROL ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_ROL(memAddr, memMode, 3, 7)

	//-------------------------------------------------- ROR --------------------------------------------------//

	case 0x6A: // Instruction ROR (  accumulator )
		opc_ROR_A(1, 2)

	case 0x66: // Instruction ROR ( zeropage )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Zeropage(PC + 1)
		}
		opc_ROR(memAddr, memMode, 2, 5)

	case 0x76: // Instruction ROR ( zeropage,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_ZeropageX(PC + 1)
		}
		opc_ROR(memAddr, memMode, 2, 6)

	case 0x6E: // Instruction ROR ( absolute )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_Absolute(PC + 1)
		}
		opc_ROR(memAddr, memMode, 3, 6)

	case 0x7E: // Instruction ROR ( absolute,X )
		if opc_cycle_count == 1 {
			memAddr, memMode = addr_mode_AbsoluteX(PC + 1)
		}
		opc_ROR(memAddr, memMode, 3, 7)

	//-------------------------------------------------- ISB? FF --------------------------------------------------//

	// ISB (INC FOLLOWED BY SBC - IMPLEMENT IT!!!!!!)
	// FF (Filled ROM)
	case 0xFF:
		// Atari 2600 interpreter mode
		if CPU_MODE == 0 {
			// 	if Debug {
			// 		fmt.Printf("\tOpcode %02X [1 byte]\tFilled ROM.\tPC incremented.\n", opcode)
			//
			// 		// Collect data for debug interface just on first cycle
			// 		if opc_cycle_count == 1 {
			// 			debug_opc_text		= fmt.Sprintf("%04x     ISB*     ;%d", PC, opc_cycles)
			// 			dbg_opc_bytes		= bytes
			// 			dbg_opc_opcode		= opcode
			// 		}
			// 	}
			// 	PC +=1
			fmt.Printf("\tOpcode 0xFF NOT IMPLEMENTED YET!! Exiting.\n")
			os.Exit(0)

			// 6502/6507 interpreter mode
		} else {
			// fmt.Println(Memory[0x20], Memory[0x21], Memory[0x22])
			fmt.Println("Opcode 0xFF in 6507 mode. Exiting.")
			os.Exit(0)
		}

	//------------------------------------------- Unnoficial Opcodes ------------------------------------------//

	// case 0x27: // Instruction RLA ( zeropage )
	// 	if opc_cycle_count == 1 {
	// 		memAddr, memMode = addr_mode_Zeropage(PC + 1)
	// 	}
	// 	opc_U_RLA(memAddr, memMode, 2, 5)

	// case 0x64: // Instruction NOP ( zeropage )
	// 	if opc_cycle_count == 1 {
	// 		memAddr, memMode = addr_mode_Zeropage(PC + 1)
	// 	}
	// 	opc_U_NOP(memAddr, memMode, 2, 3)

	//-------------------------------------------- No Opcode Found --------------------------------------------//

	default:
		fmt.Printf("\tOPCODE %02X NOT IMPLEMENTED!\n\n", opcode)
		os.Exit(0)
	}

	// Increment Cycle
	counter_F_Cycle++

	// ---------------------- Tests ---------------------- //

	// The B flag tester
	if P[5] != 1 {
		fmt.Println("Someone tryed to change P[5] to zero. Exiting!")
		os.Exit(2)
	}

	Pause_addr := 0x35a2

	// Pause
	if PC > uint16(Pause_addr-20) {
		Debug = true
	}

	// Pause
	if PC == uint16(Pause_addr) {
		Pause = true
	}

}
