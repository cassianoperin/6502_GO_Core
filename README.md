# 6502 CPU Emulator

MOS Technology 6502 8-bit microprocessor core to be uses as a go module.

All opcodes and memory modes implemented and tested in Klaus Dormann test suite.

Currently project using this core:

https://github.com/cassianoperin/6502_console

https://github.com/cassianoperin/Atari2600


## Emulation Status

* ![100%](https://progress-bar.dev/100) [Klaus Dormann 6502 Functional test suite](https://github.com/Klaus2m5/6502_65C02_functional_tests)
* ![100%](https://progress-bar.dev/100) 56 Instructions (opcodes)
* ![100%](https://progress-bar.dev/100) 13 Memory Addressing Modes
* ![100%](https://progress-bar.dev/100) One 8-bit accumulator register (A)
* ![100%](https://progress-bar.dev/100) Two 8-bit index registers (X and Y)
* ![100%](https://progress-bar.dev/100) Seven 1-bit processor status flag bits (P)
* ![100%](https://progress-bar.dev/100) Opcode cycles counter
* ![100%](https://progress-bar.dev/100) Address BUS
* ![100%](https://progress-bar.dev/100) Data BUS
* ![100%](https://progress-bar.dev/0)   IRQs
* ![100%](https://progress-bar.dev/0)   NMIs

## Improvements

* ![100%](https://progress-bar.dev/0)   Put opcodes debug message on start OR end of opcodes

## Usage

### Import 6502 CPU into the code

`import CPU_6502 "github.com/cassianoperin/6502_GO_Core"`

### Use the functions provided

#### Initialize the CPU 

`CPU_6502.Initialize()`

#### Initialize Timers

`CPU_6502.InitializeTimers()`

#### Read ROM to the memory

`CPU_6502.ReadROM(<filename string>)`
        
#### Reset Vector: 0xFFFC | 0xFFFD (Little Endian)

`CPU_6502.Reset()`

#### Interpreter

`CPU_6502.CPU_Interpreter()`


## Documentation:

### 6502

http://www.6502.org/

https://web.archive.org/web/20150217073759/http://homepage.ntlworld.com/cyborgsystems

http://datasheets.chipdb.org/Synertek/6502.pdf

### Architecture

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

### Stack:

https://wiki.nesdev.com/w/index.php/Stack

### Interrupts:

http://www.cs.jhu.edu/~phi/csf/slides/lecture-6502-interrupt.pdf


### Buses

http://www.plingboot.com/2015/10/homebrew-6502-part-2/

https://slideplayer.com/slide/3944506/
