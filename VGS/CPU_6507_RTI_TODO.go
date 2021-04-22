package VGS

// The B flag REMINDER

// Two instructions (PLP and RTI) pull a byte from the stack and set all the flags. They ignore bits 5 and 4.
//
// The only way for an IRQ handler to distinguish /IRQ from BRK is to read the flags byte from the stack and test bit 4. The slowness of this is one reason why BRK wasn't used as a syscall mechanism. Instead, it was more often used to trigger a patching mechanism that hung off the /IRQ vector: a single byte in PROM, UVEPROM, flash, etc. would be forced to 0, and the IRQ handler would pick something to do instead based on the program counter.
//
// Unlike bits 5 and 4, bit 3 actually exists in P, even though it doesn't affect the ALU operation on the 2A03 or 2A07 CPU the way it does in MOS Technology's own chips.
