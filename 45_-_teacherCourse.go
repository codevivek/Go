package main

import (
	"fmt"
)

type teacher struct {
	firstName string
	lastName  string
	field       string
}

func (t teacher) fullName() string {
	return fmt.Sprintf("%s %s", t.firstName, t.lastName)
}

type course struct {
	title   string
	content string
	teacher
}

func (c course) details() {
	fmt.Println("Title: ", c.title)
	fmt.Println("Content: ", c.content)
	fmt.Println("Author: ", c.fullName())
	fmt.Println("Bio: ", c.field)
}

func main() {
	teacher1 := teacher{
		"Dara",
		"Singh",
		"Programming",
	}
	course1 := course{
		" Golang",
		"composition & concurrency & webdevelop.",
		teacher1,
	}
	course1.details()
}
