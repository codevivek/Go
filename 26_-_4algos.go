//36 –	Create four algorithms, each displaying the corresponding one of the
//following sequences:
/*
a)		5	10	15	20	25	30	35	40
b)		3	5	7	9	11	13	15
c)		80	70	60	50	40	30	20
d)		1	2	6	24	120	720
.*/

package main

import (
	"fmt"
)

func main(){
	i :=5
	for i<=40 {
		fmt.Print(i , " " )
		i += 5
	}

    fmt.Println()
	i =3
	for i<=15 {
		fmt.Print(i , " " )
		i += 2
	}

	fmt.Println()
	i =80
	for i>=20 {
		fmt.Print(i , " " )
		i -= 10
	}
//1	2	6	24	120	720
	fmt.Println()
	i =1
	result:=1
	for i<=6 {
		result = result*i
		fmt.Print(result , " " )
		i+=1
	}
}



