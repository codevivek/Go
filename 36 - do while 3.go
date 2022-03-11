//example do while in go
//Write an algorithm that reads an integer and determines
//whether it is even or odd.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myNumber string
	for {

		fmt.Println("Enter a number between 0 and 20 or type exit: ")
		fmt.Scanln(&myNumber)
		if _, err := strconv.ParseInt(myNumber, 10, 64); err == nil {
			fmt.Println(myNumber + " is your number.")

			myNumber, err := strconv.Atoi(myNumber)

			if err ==nil && myNumber<=20 && myNumber>=0{
			   if myNumber%2 == 0 {
				   fmt.Println(strconv.Itoa(myNumber) + " is even ")
			} else {
				   fmt.Println(strconv.Itoa(myNumber) + " is odd ")
			       }
			}else{
				fmt.Println("your number must be between 0 and 20")
			}
		}
        if myNumber =="exit"{
			return
		}
	}
	fmt.Println("bye bye , see you soon !")
}