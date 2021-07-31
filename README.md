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
* ![100%](https://progress-bar.dev/100) Opcode cycles counter
* ![100%](https://progress-bar.dev/100) Address BUS
* ![100%](https://progress-bar.dev/100) Data BUS
* ![100%](https://progress-bar.dev/100) Console Mode


## Improvements
* ![100%](https://progress-bar.dev/0) Update the core to a GO Module

## Improvements (Later)
* ![100%](https://progress-bar.dev/0) IRQs
* ![100%](https://progress-bar.dev/0) Put opcodes debug message on start OR end of opcodes
* ![100%](https://progress-bar.dev/0) CONSOLE: Interrupt loops with CTRL-C (replace hard limits)
* ![100%](https://progress-bar.dev/0) CONSOLE: Mem Dump suppress repeated lines
* ![100%](https://progress-bar.dev/0) CONSOLE: Multiple commands with ";"
* ![100%](https://progress-bar.dev/0) Opcode cycle precision mode (what is done in each cycle)

## EMULATOR Build Instructions

1) MAC
* Install GO:

	 `brew install go`

* Install library requisites:

	`go get github.com/cassianoperin/pseudo-terminal-go/terminal`


* Compile:

	`go build -ldflags="-s -w" 6502.go`

2) Windows and Linux

* Not tested yet.


## Usage

	$./6502 [options] PROGRAM_NAME


- Options:

	`-debug`       Enable Debug Mode

	`-register_PC` Set the Program Counter Address (Hexadecimal)

	`-help`        Show this menu


## CONSOLE

```
Options:

	quit						Quit console

	help						Print help menu

	reset						Reinitialize CPU and reload program to memory

	step						Execute current opcode

	step <value>					Execute <value> opcodes

	step_limit <value>				Define the maximum steps allowed

	step_debug_start <value>			Set the cycle to step start showing the debug messages

	add_breakpoint <PC|A|X|Y|CYCLE>=<value>		Quit console

	del_breakpoint <index>				Delete a breakpoint

	show_breakpoints				Show breakpoints

	registers  <PC|A|X|Y>=<Value>			Change registers values

	processor_status  <N|V|B|D|I|Z|C>=<Value>	Change processor status registers values

	debug  <on|off>					Enable or Disable Debug mode

	mem						Dump full memory

	mem <address>					Dump memory address
 
 	mem <start address> <end address>		Dump memory address range
 
 	disassemble <address>				Disassemble memory address

 	disassemble <start address> <end address>	Disassemble memory address range

 	goto <address>					Run until PC=<address>
 
 	goto_limit <value>				Define the maximum steps allowed in GOTO
 
 	run						Run the emulator
 
 	run_limit <value>				Define the maximum steps allowed in RUN
```

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





 

