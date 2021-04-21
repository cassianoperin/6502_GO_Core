# 6502
6502 / 6507 Emulator written in Go


Changelog:

- Added opcodes
CMP (CD)
PLP (28)
PHP (08)


- Improved opcodes for 6502:
PHA
PLA


- Added reset Stack Pointer (SP) to 0xFF to Reset() function



Implement stack:
https://wiki.nesdev.com/w/index.php/Stack

Change all opcodes that uses it
    PHA, PLA, PLP, PHP                               (OK)
    Correct:
    JSR - Jump to New Location Saving Return Address
    RTS - Return from Subroutine

PC on 0x0674 on 6502 test
