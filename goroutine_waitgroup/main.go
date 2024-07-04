package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mut sync.Mutex
var results =[]string{"test"}

func main() {
	var wg sync.WaitGroup

	weblist := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://go.dev",
	}

	for _, web := range weblist {
		go webrequest(web, &wg) //creates multiple go routine
		wg.Add(1)
	}
	wg.Wait() // wait for all goroutine to complete
	fmt.Println(results)
}

func webrequest(s string, wg *sync.WaitGroup) {
	defer wg.Done() // give a signal to that go routine that it is completed
	res, err := http.Get(s)
	if err != nil {
		fmt.Println("Error in fetching")
	} else {
		  mut.Lock()
          results=append(results, s)
		  mut.Unlock()
		fmt.Printf("%d is the status code for %s\n", res.StatusCode, s)
	}
}
