package main

import "fmt"

func main() {
	var x int = 114
	name := "rajesh"
	print(x)
	print(name)
}
func print(data interface{}) {

	switch val := data.(type) {
	case string:
		fmt.Printf("%s is a string\n", val)
	case int:
		fmt.Printf("%d is an integer\n", val)
	default:
		fmt.Print("unsupported format\n")
	}

	/*if val, ok := data.(string); ok {
		fmt.Printf("%s is a string\n", val)
	} else {
		fmt.Print("Unsupported type\n")
	}*/
}
