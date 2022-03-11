package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//If you don't specify type here
	var b byte = 'a'

	fmt.Println("Priting Byte:")
	//Print Size, Type and Character
	fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(b), reflect.TypeOf(b), b)

	r := '£'

	fmt.Println("\nPriting Rune:")
	//Print Size, Type, CodePoint and Character
	fmt.Printf("Size: %d\nType: %s\nUnicode CodePoint: %U\nCharacter: %c\n", unsafe.Sizeof(r), reflect.TypeOf(r), r, r)

	s := "µ" //Micro sign
	fmt.Println("\nPriting String:")
	fmt.Printf("Size: %d\nType: %s\nCharacter: %s\n", unsafe.Sizeof(s), reflect.TypeOf(s), s)


    r1 := 45
    if reflect.TypeOf(r1).Kind() == reflect.Int{
		fmt.Println("it is true")
	}




}

/*
Output:

Priting Byte:
Size: 1
Type: uint8
Character: a

Priting Rune:
Size: 4
Type: int32
Unicode CodePoint: U+00A3
Character: £

Priting String:
Size: 16
Type: string
Character: µ
Caveats
Declaring a  byte with a NON-ASCII character will raise a compiler error as below. I tried with a character having a corresponding code as 285
constant 285 overflows byte
Only a single character can be declared inside a single quote while initializing byte or a rune. On trying to add two character between single quote, below compiler warning will be generated
invalid character literal (more than one character)
*/