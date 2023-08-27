package main

import (
	"fmt"
	"sync"
)

func main() {
	const gr = 500

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var m sync.Mutex //we can use eiter Mutex or Channels
	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock() // Lock() will lock the access to this code(in our case the varibale n) to other goroutines(so only one goroutine can access it at a time), until the Unlock() is called
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			m.Lock()
			defer m.Unlock() //this will execute just before the function will exit(convinient way to write)
			n--
			//m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n)
}

//race detector tool in go will help to detect the data race
//go run -race dataRace.go //will output if there's any datarace with details
