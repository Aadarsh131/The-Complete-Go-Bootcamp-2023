package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func checkURLandWriteFile(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close() //make sures there is no data rescource leak, also it will exit before the our function exits
		fmt.Println("Url: ", url, " StatusCode: ", resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			err = os.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}

func main() {
	urls := []string{"https://www.google.com", "https://www.medium.com", "https://www.facebook.com"}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkURLandWriteFile(url, &wg)
		fmt.Println(strings.Repeat("=", 50))
	}
	//fmt.Println("No. of goroutines present", runtime.NumGoroutine())
	//wg.Wait()
	//fmt.Println("No. of goroutines present", runtime.NumGoroutine()) //why the goroutines are increased?
}
