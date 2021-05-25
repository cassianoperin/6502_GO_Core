# 6502 Emulator

MOS Technology 6502 8-bit microprocessor emulator written in Go.

All opcodes and memory modes implemented and tested in Klaus Dormann test suite.


## Emulation Status
* ![100%](https://progress-bar.dev/100) 56 Instructions (opcodes)
* ![100%](https://progress-bar.dev/100) 13 Memory Addressing Modes
* ![100%](https://progress-bar.dev/100) 8-bit accumulator register (A), two 8-bit index registers (X and Y), 7 1-bit processor status flag bits (P)
* ![100%](https://progress-bar.dev/100) [Klaus Dormann 6502 Functional test suite](https://github.com/Klaus2m5/6502_65C02_functional_tests)

## TODO

* ![100%](https://progress-bar.dev/0) Review all debug messages and show in Hexadecimal
* ![100%](https://progress-bar.dev/0) Invert bytes from opcodes
* ![100%](https://progress-bar.dev/0) Improve debug functions to detect 1, 2 or 3 bytes opcodes
* ![100%](https://progress-bar.dev/0) Create a function for extra cycle detection
* ![100%](https://progress-bar.dev/0) Create a disassembler mode
* ![100%](https://progress-bar.dev/0) Create a CLI
* ![100%](https://progress-bar.dev/0) Implement a BUS from CPU to Memory (address and data)


## Documentation:

### 6502:

http://www.6502.org/

https://web.archive.org/web/20150217073759/http://homepage.ntlworld.com/cyborgsystems

## Architecture

http://www.obelisk.me.uk/6502/architecture.html

http://www.weihenstephan.org/~michaste/pagetable/6502/6502.jpg

https://www.bbc.co.uk/bitesize/guides/zr8kt39/revision/4


### Opcodes:

https://www.masswerk.at/6502/6502_instruction_set.html

http://www.obelisk.me.uk/6502/reference.html

### Addressing:

https://slark.me/c64-downloads/6502-addressing-modes.pdf

http://www.obelisk.me.uk/6502/addressing.html

http://www.emulator101.com/6502-addressing-modes.html

## Stack:

https://wiki.nesdev.com/w/index.php/Stack

## Interrupts:

http://www.cs.jhu.edu/~phi/csf/slides/lecture-6502-interrupt.pdf

## Unofficial opcodes:

https://wiki.nesdev.com/w/index.php/Programming_with_unofficial_opcodes


# REVIEW

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

https://www.bbc.co.uk/bitesize/guides/zr8kt39/revision/4





 

