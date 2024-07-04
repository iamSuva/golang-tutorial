package main

import (
	"fmt"
	"math"
)
type Shape interface{
	area() float64
	circumtance() float64
}
type square struct{
     length float64
}
type circle struct{
	radius float64
}
func (s square) area() float64{
	return s.length*s.length
}
func (s square) circumtance() float64{
	return 4*s.length
}
func(c circle) area()float64{
	return math.Pi * c.radius *c.radius
}
func (c circle) circumtance() float64{
	return math.Pi*c.radius*c.radius
}
func getShpe(s Shape){
	fmt.Printf("the area of shape is %v\n",s.area())
	fmt.Printf("the circum of shape is %v\n",s.circumtance())
}
func main(){

	fmt.Println("welcome to interfaces")

	shapes :=[]Shape{circle{4},square{5}};
	for _, v:=range shapes{
		// getShpe(v)
		a:=v.area()
		cir:=v.circumtance()
		fmt.Println("area is ",a);
		fmt.Println("circumtances is ",cir)
		fmt.Println(".......")
	}





	


}