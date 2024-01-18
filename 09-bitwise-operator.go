package main

import "fmt"
/* 
	&  bitwise AND
	|  bitwise OR
	^  bitwise XOR
	&^ bit clear (AND NOT)
	<< left shift
	>> right shift
*/
var x uint8 = 1 << 1 | 1 << 5
var y uint8 = 1 << 1 | 1 << 2

fmt.Printf("%08b\n", x)     // "00100010", the set {1, 5}
fmt.Printf("%08b\n", y)     // "00000110", the set {1, 2}

fmt.Printf("%08b\n", x&y)   // "00000010", the intersection { 1 }
fmt.Printf("%08b\n", x|y)   // "001000110", the union {1, 2, 5}

fmt.Printf("%08b\n", x^y)   // "00100100", the symmetric difference {2, 5}
fmt.Printf("%08b\n", x&^y)  // "001100000", the diference { 5 }

for i := uint(0); i < 8; i++ {
	if x&(1<<i) != 0 {
		// membership test
		fmt.Println(i) // 1, 5
	}
} 

fmt.Printf("%08b\n", x<<1)  // "01000100", the set {2, 6}
fmt.Printf("%08b\n", x>>1)  // "00010001", the set {0, 4}