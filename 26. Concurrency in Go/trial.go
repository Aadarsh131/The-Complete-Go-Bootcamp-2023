package main

import (
	"fmt"
)

func foo1(c chan int) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		//fmt.Println("f1, i=", i)
		c <- i
		//time.Sleep(time.Second) //simulating an expensive task
	}
	//fmt.Println("f1 execution finished")
}
func foo2() {
	fmt.Println("f2 normal function execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")
}
func main() {
	c := make(chan int)

	fmt.Println("Main executing started")
	go foo1(c)
	//fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	foo2()
	//time.Sleep(time.Second * 3) //giving some time, for our created goroutine to run, until the main goroutine is paused(paused for 3 sec in our case) [NOT at all recommended]

	fmt.Println("main execution stopped")
}
