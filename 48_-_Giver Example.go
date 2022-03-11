//reference:https://stackoverflow.com/questions/26027350/go-interface-fields
package main
import "fmt"

// Interface
type Giver interface {
Give() int64
}

// One implementation
type FiveGiver struct {}

func (fg *FiveGiver) Give() int64 {
return 5
}

// Another implementation
type VarGiver struct {
number int64
}

func (vg *VarGiver) Give() int64 {
return vg.number
}

// A function that uses the interface
func GetSomething(aGiver Giver) {
	fmt.Println("The Giver gives: ", aGiver.Give())
}

// Bring it all together
func main() {
	fg := &FiveGiver{}
	vg := &VarGiver{3}
	GetSomething(fg)
	GetSomething(vg)
}