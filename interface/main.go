package main

import (
	"fmt"
	"math"
)
type Shape interface{
	area() float64
	circumtance() float64
}
type Rect struct{
	width,height float64  // width and height of rectangle
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
func (r Rect) area() float64{
	return r.height*r.width
}
func (r Rect) circumtance() float64{
	return 2*(r.width+r.height)
}
func getShpe(s Shape){
	fmt.Printf("the area of shape is %v\n",s.area())
	fmt.Printf("the circum of shape is %v\n",s.circumtance())
}
func getarea(s Shape){
	fmt.Println("the are is : ",s.area())
	fmt.Println("the circum is : ",s.circumtance())
}

func main(){

	fmt.Println("welcome to interfaces")

	// shapes :=[]Shape{circle{4},square{5}};
	// for _, v:=range shapes{
	// 	// getShpe(v)
	// 	a:=v.area()
	// 	cir:=v.circumtance()
	// 	fmt.Println("area is ",a);
	// 	fmt.Println("circumtances is ",cir)
	// 	fmt.Println(".......")
	// }

	// rect:=Rect{5,6}
	// getarea(rect); 



  a:="hii this is suva"
  b:=96
  c:=93.6


display(a) // empty interface type can take any no type of parameter
display(b)	
display(c)
displayall(a,b,c,"suvadip","rohit")
 var interf interface{};
 interf=10
 fmt.Println(interf)
 interf=float64(3.14)
 fmt.Println(interf)
 interf="jjnhg hhjhb"
 fmt.Println(interf)
}

func display(i interface{}){ //empty interface type
	fmt.Println(i)
}
func displayall(i... interface{}){
	// this is take multiple parameters
	for _, v:=range i{
        fmt.Println(v)
    }
}