package main

import "fmt"

func foo(x ...int) {
	fmt.Printf("x is of type: %T \n", x)
	fmt.Println(x)
	fmt.Println("=================")
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
