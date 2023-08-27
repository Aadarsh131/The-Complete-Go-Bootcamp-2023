/////////////////////////////////
// Unbuffered Channels
// Go Playground: https://play.golang.org/p/_44csjQDJvM
/////////////////////////////////

//unbuffered channels are also sometimes called synchronous channels

package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int) //unbuffered channel

	// Launching a goroutine
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10 //as this is blocking call, any code after this in this function, will be awaited until the main function recieves the message of the channel
		fmt.Println("func goroutine after sending data into the channel")
	}(c1) // calling the anonymous func and passing c1 as argument

	//fmt.Println("main goroutine sleeps for 2 seconds")
	//time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	// we sleep for a second to give time to the goroutine to finish
	//time.Sleep(time.Second)

	// After running the program we notice that the sender (the func goroutine) blocks on the channel
	// until the receiver (the main goroutine) receives the data from the channel.
}
