package main

import (
	"fmt"
	"strconv"
)
func main() {
	if sum(2, 9) > 10 {
		fmt.Println("the sum is : " + strconv.Itoa(sum(2, 9)) + " and it is bigger than 10")
	}
}
func sum(a int, b int) int {
	var somme int = 0

	somme = a + b

	return somme

}
