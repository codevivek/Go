package main

import "fmt"

func main() {

	first := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	second := [3][2]int{{7, 8}, {9, 10}, {11, 12}}
	mult := [2][2]int{{0, 0}, {0, 0}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				mult[i][j] = mult[i][j] + (first[i][k] * second[k][j])
			}
			fmt.Print(mult[i][j], "\t")

		}
		fmt.Println()
	}

}
