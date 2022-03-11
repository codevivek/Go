package main

import (
	"fmt"
)

func main() {

	firstNames := [3]string{"dara1", "dara2", "dara3"}
	fmt.Println(firstNames)

        for i:=0;i<3;i++{
          fmt.Println(firstNames[i])

        } 


	var a [3]int
	for i := 0; i < 3; i++ {
		a[i] = i
		fmt.Println(a[i])
	}

	fmt.Println(a)
}
