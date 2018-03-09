package main

import "fmt"

func main() {

	half := func(x int) (int, bool) {
		return x /2, x % 2 == 0
	}

	fmt.Printf("%T\n", half)

	for true {
		var x int
		fmt.Print("Enter a natural number: ")
		fmt.Scanln(&x)
		y, even := half(x)
		fmt.Println("Result: ", y)
		fmt.Println("x was even: ", even)
		fmt.Println("================================")
	}
}
