# 6502 Emulator

MOS Technology 6502 8-bit microprocessor emulator written in Go.

All opcodes and memory modes implemented and tested in Klaus Dormann test suite.

## Emulation Status

* ![100%](https://progress-bar.dev/100) [Klaus Dormann 6502 Functional test suite](https://github.com/Klaus2m5/6502_65C02_functional_tests)
* ![100%](https://progress-bar.dev/100) 56 Instructions (opcodes)
* ![100%](https://progress-bar.dev/100) 13 Memory Addressing Modes
* ![100%](https://progress-bar.dev/100) One 8-bit accumulator register (A)
* ![100%](https://progress-bar.dev/100) Two 8-bit index registers (X and Y)
* ![100%](https://progress-bar.dev/100) Seven 1-bit processor status flag bits (P)
* ![100%](https://progress-bar.dev/100) Opcode cycle count
* ![100%](https://progress-bar.dev/100) Address BUS
* ![100%](https://progress-bar.dev/100) Data BUS
* ![100%](https://progress-bar.dev/70) Console Mode


## Improvements

* ![100%](https://progress-bar.dev/0) CONSOLE: Add GOTO function
* ![100%](https://progress-bar.dev/0) CONSOLE: Add Reset function
* ![100%](https://progress-bar.dev/50) CONSOLE: Create a disassembler to just decode address ranges
* ![100%](https://progress-bar.dev/0) CONSOLE: Add a Memory position or value Breakpoint
* ![100%](https://progress-bar.dev/0) CONSOLE: debug skip for x cycles to increase debug speed OR cycle to enable debug on RUN?
* ![100%](https://progress-bar.dev/0) CONSOLE: Add memory representation of current memory mode
* ![100%](https://progress-bar.dev/0) CONSOLE: Change registers, PC and memory values
* ![100%](https://progress-bar.dev/0) Other Klaus tests
* ![100%](https://progress-bar.dev/0) Review Console and CLI codes
* ![100%](https://progress-bar.dev/0) Put debug message on start OR end of opcodes
* ![100%](https://progress-bar.dev/0) IRQs
* ![100%](https://progress-bar.dev/0) Opcode cycle precision mode





## Documentation:

### 6502

http://www.6502.org/

https://web.archive.org/web/20150217073759/http://homepage.ntlworld.com/cyborgsystems

http://datasheets.chipdb.org/Synertek/6502.pdf

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


## BUSES

http://www.plingboot.com/2015/10/homebrew-6502-part-2/

https://slideplayer.com/slide/3944506/

## Console

https://tutorialedge.net/golang/reading-console-input-golang/





 

