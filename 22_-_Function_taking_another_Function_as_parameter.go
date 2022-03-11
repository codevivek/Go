package main

import "fmt"

type fn func(string)

func firstFunction(s1 string) {
	fmt.Printf("\ns1 is %v", s1)
}
func secondFunction(s2 string) {
	fmt.Printf("\ns2 is %v", s2)
}
func tryFn(f fn, val string) {
	f(val)
}
func main() {
	tryFn(firstFunction, "from first function")
	tryFn(secondFunction, "from second function")
}
