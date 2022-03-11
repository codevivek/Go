//Go’s structs are typed collections of fields.
//They’re useful for grouping data together to form records.
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//This person struct type has name and age fields.
type person struct {   //This syntax creates a new struct.
	name string
	age  int
}
//NewPerson constructs a new person struct with the given name
//It’s idiomatic to encapsulate new struct creation in constructor functions
func NewPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p    //You can safely return a pointer to local variable as a local variable will survive the scope of the function.
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"}) //Omitted fields will be zero-valued.
	fmt.Println(&person{name: "Ann", age: 40})  //An & prefix yields a pointer to the struct.
	fmt.Println(NewPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)    //Access struct fields with a dot.

	sp := &s
	fmt.Println(sp.age)      //You can also use dots with struct pointers - the pointers are automatically dereferenced.

	sp.age = 51        //Structs are mutable.
	fmt.Println(sp.age)

	fmt.Println(&person{name: "Ann", age: 40})  //An & prefix yields a pointer to the struct.
	s = person{name: "Sean", age: 50}
	fmt.Println(s.age)    //Access struct fields with a dot.

	fmt.Println("--------")
	fmt.Println(s)   //this contains the content of s
	fmt.Printf("Size: %d\nType: %s\n", unsafe.Sizeof(s), reflect.TypeOf(s))

	fmt.Println("--------")
	sp = &s
	fmt.Println(sp)  //this contains the address of s
	fmt.Printf("Size: %d\nType: %s\n", unsafe.Sizeof(sp), reflect.TypeOf(sp))

	fmt.Println(sp.age)      //You can also use dots with struct pointers - the pointers are automatically dereferenced.




}