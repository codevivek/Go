package main

	import (
		"fmt"
	)

	func main() {

		var coins1, coins2 bool
		fmt.Println("question 1: ")
		fmt.Scan( &coins1) //
		fmt.Println("question 2: ")
		fmt.Scan( &coins2)
		//fmt.Println("question 3")
		fmt.Scan( &coins2)
		fmt.Print(coins1, coins2)

	}