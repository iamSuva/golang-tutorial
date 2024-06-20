package main

//  a method is a 
// function that is
//  defined on struct types.

import "fmt"
func main(){
	fmt.Println("methods in golang")
    rect:=shape{5,8};
	fmt.Print("area of rectangle is ",rect.area());
}

type shape struct{
	x int
	y int
}

func (s shape) area()int{
	return s.x *s.y;
}