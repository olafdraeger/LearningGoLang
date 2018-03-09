package main

import (
	"fmt"
	"math"
)

func findmax(x ...int) int {
	fmt.Printf("x is of type: %T \n", x)
	var max float64
	for _, m := range x {
		max = math.Max(float64(m), max)
	}
	return int(max)
}
func main() {
	fmt.Println(findmax(1, 2, 3, 4, 5, 23, 4, 42, 22, 54, 65, 99, 21,55))
}
