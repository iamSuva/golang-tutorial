package main

import (
	"fmt"
	// "time"
)
func sayGo(){
	fmt.Println("say go ")
	//time.Sleep(1000*time.Millisecond)
	fmt.Println("waiting is over")
   
}
func sayHello(){
	fmt.Println("say hello:)");
}
func main() {
	fmt.Println("this is go routine function")
	go sayGo()
	go sayHello()
	//time.Sleep(900*time.Millisecond)
}
