/*
Example 2: Composition through Embedding in Interfaces

In the Go language interfaces are implicitly implemented. That is to say,
if methods, which are defined in an interface, are used on objects such as
structs, then the struct is said to implement the interface. An interface
can be embedded with other interfaces in order to form composite interfaces.
If all the interfaces in a composite interface are implemented, then the
composite interface is also said to be implemented by that object.
*/
// Golang Program to implement composite interfaces
package main

import "fmt"

type purchase interface {
	sell()
}

type display interface {
	show()
}

// this is a composite interface
type salesman interface {
	purchase
	display
}

type book struct {
	name, price string
	bookCollection []string
}

// We use the book struct to
// implement the interfaces
func (b book) sell() {
	fmt.Println("--------------------------------------")
	fmt.Println("Name:", b.name)
	fmt.Println("Price:", b.price)
	fmt.Println("--------------------------------------")
}
func (b book) show() {
	fmt.Println("The books available are: ")
	for _, name := range b.bookCollection {
		fmt.Println(name)
	}
	fmt.Println("--------------------------------------")
}

// This method takes the composite
// interface as a parameter
// Since the interface is composed
// of purchase and display
// Hence the child methods of those
// interfaces can be accessed here
func shop(s salesman) {
	fmt.Println(s)
	s.sell()
	s.show()
}

func main() {

	collection := []string{"JavaScript",
		"C++", "Ruby"}
	book1 := book{"Coding by example", "$55", collection}
	shop(book1)

}

