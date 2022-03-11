//14 –	A print shop charges 5 cents per copy for the first 100 copies.
//For any subsequent copies, they charge 3 cents each. Write an algorithm
//that determines the price associated with a given number of copies.

package main

import (
	"fmt"
)

func main() {
	for {

		var numberOfCopy float32 = 0
		fmt.Println("Enter the number of copies: ")
		fmt.Scanln(&numberOfCopy)

		const limit float32 = 100
		const lowPrice float32 = 3
		const highPrice float32 = 5

		if numberOfCopy <= limit {
			var bill float32 = float32(numberOfCopy) * highPrice / limit
			s := fmt.Sprintf("%.2f", bill) //
			fmt.Println("the bill is : " + s)
		}

		if numberOfCopy > limit {
			var bill float32 = ((limit) * highPrice / limit) + ((numberOfCopy - limit) * lowPrice / limit)
			s := fmt.Sprintf("%.2f", bill) //
			fmt.Println("the bill is : " + s + "$")
		}

		check := 0
		fmt.Println("Enter 1 to continue or Enter exit: ")
		n, err := fmt.Scanln(&check)
		//fmt.Println(err)
		if n < 1 || err != nil {
			fmt.Println("bye bye , see you soon !")
			return
		}
	}

}
