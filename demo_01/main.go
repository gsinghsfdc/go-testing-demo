package main

import "fmt"

func main() {
	fmt.Printf("Sum of %v and %v is %v", 2, 4, AddTwoNumbers(2, 4))
}

func AddTwoNumbers(a, b int) int {
	return a + b
}
