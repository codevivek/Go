package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a [3][2][2]int
	for i := 0; i < 3; i++ {
		fmt.Println()
		for j := 0; j < 2; j++ {
			fmt.Println()
			for k := 0; k < 2; k++ {
				a[i][j][k] = i + j + k
				if k == 0 {
					fmt.Print(strconv.Itoa(a[i][j][k]) + ",")
				} else {
					fmt.Print(strconv.Itoa(a[i][j][k]))
				}
			}
		}
	}

}
