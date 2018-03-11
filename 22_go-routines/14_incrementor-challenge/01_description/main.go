package main

import (
	"fmt"
	//"sync"
	//"sync/atomic"
	"sync"
	//"reflect"
)

var count int
//var wg sync.WaitGroup

func main() {
	//wg.Add(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go incrementor("1", ch1)
	go incrementor("2", ch2)
	//wg.Wait()
	ch := merge(ch1, ch2)
	for n := range ch {
		count += n
	}
	fmt.Println("Final Counter:", count)
}

func incrementor(s string, ch chan int) {
	for i := 0; i < 20; i++ {
		//atomic.AddInt64(&count, 1)
		ch <- 1
		fmt.Println("Process: "+s+" printing:", i)
	}
	close(ch)
	//wg.Done()
}

func merge(cs ...chan int) <- chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <- chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	// range spits our two values, index and chan
	for _, ch := range cs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
