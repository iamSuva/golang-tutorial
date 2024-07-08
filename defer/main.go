package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("this is log message")
	// fmt.Println("Defer to be executed at the end")
	// defer fmt.Println("this is defer1")
	// defer fmt.Println("this is defer2")
	// defer fmt.Println("this is defer3")
	// fmt.Println("hello world")
	// defer fmt.Println("after panic is occurred")
	// fmt.Println("Hello this is suvadip")
	// panic("hii some errored occured")
	// fmt.Println("end of execution")

	division(20, 5)
	division(24, 0)
	division(62, 4)
}

func HandlePanic() {
	//The recover statement prevents the termination of the program and recovers the program from panic.
	r := recover()
	// detect if panic occurs or not
	// we have used r := recover() to detect any
	// occurrence of panic in our program and assign the panic message to a.
	if r != nil {
		fmt.Println("recoverd from : ", r)
	}
}
func division(a, b int) {
	defer HandlePanic() //execute the handlePanic even after panic occurs
	if b == 0 {
		panic("divisor cannot be zero")

	}
	fmt.Println("result is : ", a/b)

	/*
	We are calling the handlePanic() before the occurrence of panic.
	It's because the program will be terminated if it encounters panic
	and we want to execute handlePanic() before the termination.
	We are using defer to call handlePanic().
	It's because we only want to handle the panic after it occurs,
	so we are deferring its execution.
	*/
}

//lifo order -> 3 2 1
