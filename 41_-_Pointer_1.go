package main
import 	"fmt"

func main() {
	a := 20
	fmt.Println(a)
	fmt.Println(&a) // &a is a pointer to a memory address.
	fmt.Println(*(&a)) //dereferencing
	fmt.Printf("%T\n", &a)
    fmt.Println("--------------------")
	var b = &a
	fmt.Println(b ) //  b is a pointer to the memory address where an int is stored

	// b is of type 'int pointer'.
	// *int -- the * is part of the type -- b is of type *int
	fmt.Printf("%T\n", b)

	// *b is known as derefering. the * is an operator in this case
	fmt.Println(*b)
	fmt.Printf("%T\n", *b)

	*b = 40        // b says, "The value at this address, change it to 40"
	fmt.Println(a) // 40
}
