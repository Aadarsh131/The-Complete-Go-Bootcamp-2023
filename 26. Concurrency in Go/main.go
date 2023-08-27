package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second) //simulating an expensive task
	}
	fmt.Println("f1 execution finished")
	wg.Done()
	//OR
	//(*wg).Done()
}
func f2() {
	fmt.Println("f2 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")
}
func main() {
	fmt.Println("Main executing started")
	//fmt.Println("No. of CPUs:", runtime.NumCPU())
	//fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	//fmt.Println("OS:", runtime.GOOS)
	//fmt.Println("Arch:", runtime.GOARCH)
	//fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup
	wg.Add(1) //argument tells the number of goroutines to wait for

	go f1(&wg) //if I dont dont use waitgroup and do a Wait() iside the main thread, then, main function will exit without executing anytning inside the child goroutine , we can also use a channel
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	f2()
	//time.Sleep(time.Second * 3) //giving some time, for our created goroutine to run, until the main goroutine is paused(paused for 3 sec in our case) [NOT at all recommended]

	wg.Wait() //stopping the main goroutine until the specified child goroutine has finished its  execution
	fmt.Println("main execution stopped")
}
