package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the Number:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.Trim(input, "[]")
	num := strings.Split(input, ",")
	for _, s := range num {
		s = strings.TrimSpace(s)
		nums, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid number")
			continue
		}
		if nums%2 == 0 {
			fmt.Printf("the num %v is even \n", nums)
		} else {
			fmt.Printf("the num %v is odd \n", nums)
		}
	}

}

/*func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println("The Number is Even")
	} else {
		fmt.Println("The Number is ODD")
	}
}*/
