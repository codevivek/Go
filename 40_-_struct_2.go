package main
import "fmt"

type Vertex struct {
	X, Y int
}

var ( //global variables
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex

	a string = "dara"
)
func foo1(){
fmt.Println(a + " inside foo1")
}

func foo2(){
fmt.Println(a + " inside foo2")
}
func main() {
	fmt.Println(v1, *p,p, v2, v3)
	foo1()
	foo2()
}

