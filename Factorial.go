package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1

	} else {
		return num * factorial(num-1)
	}
}
func main() {
	var n int
	fmt.Print("Enter a Number")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Print("The Number is NOt Defined Properly")
	} else {
		res := factorial(n)
		fmt.Printf("the factorial of %v is %v", n, res)
	}
}
