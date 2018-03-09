package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println("the address of the variable a as standard address is: ",&a)
	doit(&a)
	fmt.Println("back in main:", a)
}

func doit(x *int) {
	fmt.Println("\naddress in doit now: ", x)
	fmt.Println("value:",*x)
	*x++
	fmt.Println("value post incremebt:",*x)
}