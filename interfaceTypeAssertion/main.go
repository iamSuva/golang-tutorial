package main

import "fmt"

func main() {
	var a interface{}
	a = 10
	interfaceval := a.(int)
	fmt.Println("value is ", interfaceval)
	a="suvadip"
	interfaceval2,boolval:=a.(int) // as a is type of string 
	fmt.Println("value is ", interfaceval2," bool value is ",boolval)
}
