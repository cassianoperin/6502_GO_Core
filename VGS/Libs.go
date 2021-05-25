package VGS

import "fmt"

// ---------------------------- Library Function ---------------------------- //

// Memory Page Boundary cross detection
func MemPageBoundary(Address1, Address2 uint16) bool {

	var cross bool = false

	// Get the High byte only to compare
	// Page Boundary Cross detected
	if Address1>>8 != Address2>>8 {
		cross = true

		if Debug {
			fmt.Printf("\tMemory Page Boundary Cross detected! Add 1 cycle.\tPC High byte: %02X\tBranch High byte: %02X\n", Address1>>8, Address2>>8)
		}
		// NO Page Boundary Cross detected
	} else {
		if Debug {
			fmt.Printf("\tNo Memory Page Boundary Cross detected.\tPC High byte: %02X\tBranch High byte: %02X\n", Address1>>8, Address2>>8)
		}
	}

	return cross
}

// Decode Two's Complement
func DecodeTwoComplement(num byte) int8 {

	var sum int8 = 0

	for i := 0; i < 8; i++ {
		// Sum each bit and sum the value of the bit power of i (<<i)
		sum += (int8(num) >> i & 0x01) << i
	}

	return sum
}

// // BCD - Binary Coded Decimal
// func BCD(number byte) byte {

// 	var tmp_hundreds, tmp_tens, tmp_ones, bcd byte

// 	// Split the Decimal Value
// 	tmp_hundreds = number / 100    // Hundreds
// 	tmp_tens = (number / 10) % 10  // Tens
// 	tmp_ones = (number % 100) % 10 // Ones

// 	fmt.Printf("H: %d\tT: %d\tO: %d\n", tmp_hundreds, tmp_tens, tmp_ones)

// 	// Combine in one decimal number
// 	bcd = (tmp_hundreds * 100) + (tmp_tens * 10) + tmp_ones

// 	return bcd
// }

// Memory Bus - Used by INC, STA, STY and STX to update memory and sinalize TIA about the actions
func memUpdate(memAddr uint16, value byte) {

	Memory[memAddr] = value

}
