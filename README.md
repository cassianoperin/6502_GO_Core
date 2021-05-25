# 6502
6502 / 6507 Emulator written in Go

# TODO

## Show everything in hex mode? Put 0x on hex?

## Invert bytes from from opcodes (all opcodes with 3 bytes) -> Improve to detect and show correctly 2 or 3 without the if (create a string?)

## create a function for extra cycle detection (and invert to first check the opcode and later the cycle?)

## CLI

https://wiki.nesdev.com/w/index.php/Programming_with_unofficial_opcodes

## JMP Indirect - Handle addresses with FF

http://6502.org/tutorials/6502opcodes.html

https://www.reddit.com/r/EmuDev/comments/fi29ah/6502_jump_indirect_error/


JMP transfers program execution to the following address (absolute) or to the location contained in the following address (indirect). Note that there is no carry associated with the indirect jump so:

AN INDIRECT JUMP MUST NEVER USE A
VECTOR BEGINNING ON THE LAST BYTE
OF A PAGE

For example if address $3000 contains $40, $30FF contains $80, and $3100 contains $50, the result of JMP ($30FF) will be a transfer of control to $4080 rather than $5080 as you intended i.e. the 6502 took the low byte of the address from $30FF and the high byte from $3000. 

It's a bug in the 6502 that wraps around the LSB without incrementing the MSB. So instead of reading address from 0x02FF-0x0300 you should be looking at 0x02FF-0x0200. The A900 printed in the log is the value at 0x02FF-0x0300 which is not what's actually being used.

## Tests

https://github.com/Klaus2m5/6502_65C02_functional_tests/blob/master/bin_files/6502_functional_test.lst


## Addressing modes:

https://slark.me/c64-downloads/6502-addressing-modes.pdf


## Hex Calculator

https://www.calculator.net/hex-calculator.html?number1=998&c2op=%2B&number2=2&calctype=op&x=93&y=29


## Architecture

http://www.weihenstephan.org/~michaste/pagetable/6502/6502.jpg


## BUSES

http://www.plingboot.com/2015/10/homebrew-6502-part-2/

https://slideplayer.com/slide/3944506/

* SET ADDRESS and GET SOME DATA ON DATA LINES

* On reset, the processor will read address $FFFC and $FFFD (called the reset vector) and load the program counter (PC) with their content. 

* Create the 40 pins, starting with 16 address, 8 from data and the reset one

* R/W PIN
A READ operation (the R/W line is at logic one) causes eight bits of information (usually called data) to be transferred over the data bus, from the memory location specified by the address on the address bus to an 8-bit register in the microprocessor.

A WRITE operation (the RjW line is at logic zero) causes eight bits of information to be transferred from an 8-bit register in the microprocessor to a  memory location specified by the address on the address bus. The words "load" and "store" are sometimes used synonymously with the words "read" and "write," respectively. Because data are moved in one direction by a read or load opera-tion and in the other direction by a write or store operation, the data bus is said to be hidirectional. Furthermore, since data are trans-ferred as 8-bit binary numbers, that is, one byte at a time, the 6502 is called an 8-bit microprocessor.

* processor set MAR and OR receive the data of memory throug data bus OR send a value thtough data bus to that address

// P = Status Register (SR) turn into a byte
// add Instruction Register (IR) hold instruction being decoded or executed
// add Memory Data Register (MDR) holds data which has just arrived along the data bus or is just about to  be sent along data bus
// add Memory Address Register (MAR) it holds an address about to be sent out along the address bus 

## Doubts with stack:

https://wiki.nesdev.com/w/index.php/Stack

## SPLIT MEMORY ACCESS 

## Extra cycles on memory address modes? pass initial page and check if it changes?


## Basic Architecture

The only other reserved locations in the memory map are the very last 6 bytes of memory $FFFA to $FFFF which must be programmed with the addresses of the non-maskable interrupt handler ($FFFA/B), the power on reset location ($FFFC/D) and the BRK/interrupt request handler ($FFFE/F) respectively.


http://www.obelisk.me.uk/6502/architecture.html

## Opcodes

http://6502.org/tutorials/6502opcodes.html#CMP

https://www.masswerk.at/6502/6502_instruction_set.html#CLD

http://www.obelisk.me.uk/6502/reference.html

https://web.archive.org/web/20160406122905/http://homepage.ntlworld.com/cyborgsystems/CS_Main/6502/6502.htm#ADC




 

