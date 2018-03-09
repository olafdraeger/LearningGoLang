package main

import "fmt"

func main() {
	var small int
	var large int
	fmt.Print("Enter a small number: ")
	fmt.Scanln(&small)
	fmt.Print("Enter a large number: ")
	fmt.Scanln(&large)
	fmt.Printf("remainder of %d and %d is %d\n", small, large, large % small)

}