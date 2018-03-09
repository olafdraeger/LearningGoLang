package main

import (
	"fmt"
	"math/rand"
	"time"
)

//var count int = 0
func main() {
	start := time.Now()
	for x := 0; x < 10000 ;x++  {
		c := gen()
		for n := range c {
			//fmt.Println(n)
			for m := range (factorial(n)) {
				//count++
				//fmt.Println("Run No: ", count)
				fmt.Println("Factorial of: ", n, "is: ",m)
			}
		}
	}
	fmt.Println( "Factorial running time: ",time.Since(start))

}

func gen() chan int {
	out := make(chan int)
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 1; i <= 100; i++ {

			out <- r.Intn(20)
		}
		close(out)
	}()
	return out
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/
