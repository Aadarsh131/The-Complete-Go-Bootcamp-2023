package main

import (
	"fmt"
)

func foo(n int, c chan int) {
	fmt.Println("heelloow ")
	c <- n
}

func foo3() {
	fmt.Println("I have no work, just wasting my time here")
}

func foo4() {
	fmt.Println("I am from foo4")
}

func main() {

	var c chan int
	fmt.Println(c) //would be <nil> initially

	c = make(chan int) //initializes a new chan of type int (using make())
	defer close(c)     //good practice the close the channel [NOTE: defer will be executed right before the function it is in(in our case "main") ends]

	//only for recieving
	cRecieve := make(<-chan string)

	//only for sending
	cSend := make(chan<- string) //notice it is a pointer

	fmt.Printf("%T , %T , %T \n", c, cRecieve, cSend)

	go foo(10, c) //notice how the child goroutine is executed and the value is stored to the channel
	go foo3()     //notice how this function is not using any channel value, but still the goroutine is getting executed, (normally without no channel usage in our program, this child goroutine execution will not be done)

	fmt.Println(<-c) //this is a blocking call
}
