package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(800)

	var count int
	for s := range c {
		count++
		fmt.Println(s)
	}
	fmt.Println("Final Count: ", count)
}

func incrementor(n int) <- chan string{
	outC := make(chan string)
	doneB := make(chan bool)
	for i := 0; i < n; i++ {
		go func(i int) {
			for x := 0; x < 20; x++ {
				outC <- fmt.Sprint("Process: " + strconv.Itoa(i) + " printing: ", x)
			}
		doneB <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<- doneB
		}
		close(outC)
	}()
	return outC
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
