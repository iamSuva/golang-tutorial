package main

import "fmt"

func main() {
	fmt.Println("Defer to be executed at the end")
	defer fmt.Println("this is defer1")
	defer fmt.Println("this is defer2")
	defer fmt.Println("this is defer3")
	fmt.Println("hello world")
}
//lifo order -> 3 2 1