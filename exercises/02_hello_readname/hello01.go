package main

import (
	"fmt"
	//"strings"
)

func main() {
	var count int
	var myname string
	var err error
	fmt.Print("Please enter your name: ")
	count, err = fmt.Scanln(&myname)
	if err == nil {
		//strings.Join(myname, " ")
		fmt.Println("Hello " + myname + "!")
		fmt.Println(count)
	}

}
