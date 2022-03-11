package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x, y float64 = 3.45, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	var p string = strconv.Itoa(z)
	fmt.Println(x, y, f, z, p)
	fmt.Println(strconv.Itoa(z))

	q := strconv.Quote("Hello, é")
	fmt.Println(q)
	q = strconv.QuoteToASCII("Hello, é")
	fmt.Println(q)

	//float to string
	st := fmt.Sprintf("%f", 123.456)
	fmt.Println(st)

	//string to float
	f1 := "3.14159265"
	if s, err := strconv.ParseFloat(f1, 32); err == nil {
		fmt.Println(s) // 3.1415927410125732
	}
	if s, err := strconv.ParseFloat(f1, 64); err == nil {
		fmt.Println(s) // 3.14159265
	}

	v1 := "4"
	if _, err := strconv.ParseInt(v1,10,64); err == nil {
		fmt.Printf("%q looks like a number.\n", v1)
	}

	
}