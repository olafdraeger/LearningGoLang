package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 100; i++ {
		sum = sum + i
		for j := 0; j <= 100; j++ {
			fmt.Println("round:", i, "\t", j, "\t", sum+j)
		}
	}
}
