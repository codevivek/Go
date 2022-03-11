//example do while in go 
//Write an algorithm that reads an integer and determines 
//whether it is even or odd.

package main

import (
	"fmt"
	"strconv"
)

func main(){
	var myNumber int
	for  {

		fmt.Println("Enter Your number or type exit: ")
		//fmt.Scanln(&myNumber)
		//if myNumber == 0{
		//	break
		//}
		n, err := fmt.Scanln(&myNumber)
		fmt.Println(n)
		if  n<1 || err != nil {
			fmt.Println("invalid input")
			return
		}

		fmt.Println(strconv.Itoa(myNumber) + " is your number.")
		if myNumber%2 == 0 {
			fmt.Println(strconv.Itoa(myNumber) + " is even ")
		} else {
			fmt.Println(strconv.Itoa(myNumber) + " is odd ")
		}

	}
    fmt.Println("bye bye , see you soon !")
}

