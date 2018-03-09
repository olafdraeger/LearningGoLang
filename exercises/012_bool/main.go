package main

import "fmt"

func main() {
	fmt.Println((true && false))
	fmt.Println((false && true))
	fmt.Println(!(false && false))
	fmt.Println((false || true))
	fmt.Println(((true && false) || (false && true) || !(false && false)))

}
