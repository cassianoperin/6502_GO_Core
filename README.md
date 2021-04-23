# 6502
6502 / 6507 Emulator written in Go

PC on 0x095c on 6502 test (now on loop on 0x3720, needs to be fixed)



IMPLEMENT PAUSE!



# TODO

## Unofficial Opcodes

https://wiki.nesdev.com/w/index.php/Programming_with_unofficial_opcodes


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
 

