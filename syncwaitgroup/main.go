package main

import (
	"fmt"
	"sync"
)

func wokerfun(i int,wg *sync.WaitGroup) {
	defer wg.Done()
	// tell that this go routine finish execution
	fmt.Printf("worder %d started\n",i);
	fmt.Printf("worder %d ended\n",i);
}
func main() {
    var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment wait group counter
		go wokerfun(i,&wg)
	}
	wg.Wait();// wait for all workers to finish
   fmt.Println("All goroutines finished")

}