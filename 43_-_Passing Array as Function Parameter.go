package main
import "fmt"
/*
func myfun(a [10]int, size int) int {
	var k, som, resultat int

	for k = 0; k < size; k++ {
		som += a[k]
	}
    fmt.Println(som)
	resultat = som / size
	return resultat
}
*/

func myfun(a [10]int) int {
	var k, som, resultat int

	for k = 0; k < len(a); k++ {
		som += a[k]
	}
	fmt.Println(len(a))
	resultat = som / len(a)
	return resultat
}


// Main function
func main() {


	var arr = [10]int{1, 2, 3, 4, 5, 6,7,8,9,10}
	var resultat int

	resultat = myfun(arr)
	fmt.Printf("Final result is: %d ", resultat)
}