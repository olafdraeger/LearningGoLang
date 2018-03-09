package main

import "fmt"

func main() {
	f := factorial(42)
	for n := range(f){
		fmt.Println(n)
	}
}

func factorial(n int) chan uint64 {
	out := make(chan uint64)
	go func() {
		var total uint64 = 1
		for i := n; i > 0; i-- {
			total *= uint64(i)
		}
		out <- total
		close(out)
	}()

	return out
}
