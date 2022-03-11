//19 –	Give an algorithm that reads three numbers (a, b, c) and
//determines whether any one of the numbers is equal to the sum of the other two.
//If such a number exists, display it; otherwise, display the message “No solution”.

package main

import (
	"fmt"
	"strconv"
)

func main(){
	for {
		var i = 0
		var myNumbers = [3]int{}
		for i < 3 {
			fmt.Println("Please enter number  " + strconv.Itoa(i+1) + ": ")
			fmt.Scanln(&myNumbers[i])
			i++
		}

		if myNumbers[0] == myNumbers[1]+myNumbers[2] {
			fmt.Println(strconv.Itoa(myNumbers[0]) + " is : " + strconv.Itoa(myNumbers[1]) + " + " + strconv.Itoa(myNumbers[2]))
		} else if myNumbers[1] == myNumbers[0]+myNumbers[2] {
			fmt.Println(strconv.Itoa(myNumbers[1]) + " is : " + strconv.Itoa(myNumbers[0]) + " + " + strconv.Itoa(myNumbers[2]))
		} else if myNumbers[2] == myNumbers[0]+myNumbers[1] {
			fmt.Println(strconv.Itoa(myNumbers[2]) + " is : " + strconv.Itoa(myNumbers[0]) + " + " + strconv.Itoa(myNumbers[1]))
		} else {
			fmt.Println("no solution")
		}

		check  := 0
		fmt.Println("Enter 1 to continue or Enter exit: ")
		n, err := fmt.Scanln(&check)
		if  n<1 || err != nil {
			fmt.Println("bye bye , see you soon !")
			return
		}
	}
    
}

