package CORE

import "time"

var (
	CPU_MODE byte = 1 // 0 for Atari, 1 for pure 6502/6507 CPU interpreter

	// ------------------------ Hardware Components ------------------------- //
	Memory [65536]byte // Memory
	PC     uint16      // Program Counter
	A      byte        // Accumulator
	X      byte        // Index Register X
	Y      byte        // Index Register Y
	SP     byte        // Stack Pointer
	// The stack pointer is addressing 256 bytes in page 1 of memory, ie. values 00h-FFh will address memory at 0100h-01FFh.
	// As for most other CPUs, the stack pointer is decrementing when storing data.
	// However, in the 65XX world, it points to the first FREE byte on stack, so, when initializing stack to top set S=(1)FFh (rather than S=(2)00h).
	P [8]byte
	//P			byte			// Processor Status Register
	// Processor Status Register (Flags)
	// Bit  Name  Expl.
	// 7    N     Negative/Sign (0=Positive, 1=Negative)
	// 6    V     Overflow      (0=No Overflow, 1=Overflow)
	// 5    -     Not used      (Always 1)
	// 4    B     Break Flag    (0=IRQ/NMI, 1=RESET or BRK/PHP opcode)
	// 3    D     Decimal Mode  (0=Normal, 1=BCD Mode for ADC/SBC opcodes)
	// 2    I     IRQ Disable   (0=IRQ Enable, 1=IRQ Disable)
	// 1    Z     Zero          (0=Nonzero, 1=Zero)
	// 0    C     Carry         (0=No Carry, 1=Carry)

	// -------------------------------- Bus --------------------------------- //
	addressBUS uint16 // // 16 pins of processor that points to memory for read or write operations

	// --------------------------- CPU Variables ---------------------------- //
	opcode byte // CPU Operation Code

	// ------------------------------ Counters ------------------------------ //
	counter_F_Cycle uint16 // Frame Cycles
	opc_cycle_count byte   // Opcode cycle counter
	opc_cycle_extra byte   // Opcode extra cycle

	// -------------------------- Memory Variables -------------------------- //
	memMode  string // Receive the addressing mode used in the debug
	memValue int8   // Receive the memory value needed by branches. Calculated in the first opc cycle to check for extra cycles, used in the last to perform the operation

	// ------------------------------- Timers ------------------------------- //
	clock_timer *time.Ticker // CPU Clock // CPU: MOS Technology 6507 @ 1.19 MHz;
	// second_timer        = time.Tick(time.Second) // 1 second to track FPS and draws
	screenRefresh_timer *time.Ticker // Screen Refresh

	// --------------------------- Debug Interface -------------------------- //
	dbg_show_message string // Debug opcode detail messages

	// ------------------------------ Graphics ------------------------------ //
	// Screen Size
	sizeX float64 = 160.0 // 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	sizeY float64 = 192.0 // 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	// Window Resolution
	screenWidth  float64 = 1024
	screenHeight float64 = 768
	// Pixel size
	width  float64
	height float64

	// Enable or disable CPU during WSYNC
	CPU_Enabled bool

	// Pause
	Pause bool = false

	// Debug
	Debug bool = false
)
