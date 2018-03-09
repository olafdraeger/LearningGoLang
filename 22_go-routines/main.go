package main

import (
	"fmt"
)

func main() {
	var f, x uint64 = 35,1
	c := make(chan uint64)
	go func() {
		for i := f; i > 0; i-- {
			c <- i
		}
		close(c)
	}()
	//f := factorial(4)

	for n := range (c) {
		x *= uint64(n)
	}
	fmt.Println(x)
}
