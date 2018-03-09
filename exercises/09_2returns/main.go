package main

import "fmt"

func half(x int) (int, bool) {
	return x /2, x % 2 == 0
}
func main() {
	var x int
	var y int
	var even bool

	for true {
		fmt.Print("Enter a natural number: ")
		fmt.Scanln(&x)
		y, even = half(x)
		fmt.Println("Result: ", y)
		fmt.Println("x was even: ", even)
		fmt.Println("================================")
	}
}
