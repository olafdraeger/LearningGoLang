package main

import "fmt"

func main() {
	f := factorial(20)
	fmt.Println(f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
