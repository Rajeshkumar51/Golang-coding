package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}
	var p person
	p.name = "ajay"
	p.age = 12

	fmt.Print(p.name)
	fmt.Print(p.age)
}
