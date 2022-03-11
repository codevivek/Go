/*8 –	Write a program that calculates and displays the product of all
the odd numbers between 1 to 15, inclusively. */


package main

import "fmt"

func main(){

	limit :=15

	i :=1
	result:=1
	for i<=limit {

		if i%2!=0{
			result = result *i
			fmt.Println("for " , i, " : " , result )
		}
		
		i += 1
	}
      
}



