package main

import "fmt"

func init() {
	fmt.Println("ist init function called")
}
func init() {
	fmt.Println("2nd init function called")
}
func init() {
	fmt.Println("3rd init function called")
}

func main(){
	fmt.Println("this is main function")
    fmt.Println("program ended")
    
}
