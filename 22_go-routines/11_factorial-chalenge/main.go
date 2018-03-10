package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var f, count int
	start := time.Now()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 1000000; i++ {
		count++
	f := factorial(r.Intn(20))
	fmt.Println(f)
	}
	fmt.Println(f)
	elapsed := time.Since(start)
	fmt.Println("factorial runs ", count," / time :", elapsed)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
