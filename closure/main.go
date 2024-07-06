package main
import "fmt"
func MyClosure()func()int{
  sum:=0
  return func()int{
	return sum+5
  }
}
func Myfun()func(s string)string{
	str:="Hello "
	return func(s string)string{
		str+=s
		return str
	}
}
func update(v *int){
     *v=*v*15

}
func Myval(v int)*int{
	v+=32
	return &v
}
func fibo()func()int{
	a,b:=0,1
	
    return func() int {
        result :=a
		temp:=a+b
        a=b
		b=temp
        return result
    }
}
func main() {
	fmt.Println("this is closure example")
	// func(s string){
	// 	fmt.Println(s)
	// }("hello dev")
	inc:=MyClosure()
	fmt.Println(inc())
	fun:=Myfun()
	fmt.Println(fun("Suva"))
	fmt.Println(fun("dip"))
	fun2:=Myfun()
	fmt.Println(fun2("Jeet"))

	// a:=10
	// var ptr *int
	// ptr=&a
	// fmt.Println(*ptr)
	// update(&a)
	// fmt.Println(*ptr)
	// fmt.Println(a)
	// ptr:=&a
	// ptr2:=Myval(a)
	// fmt.Println(*ptr)
	// fmt.Println(*ptr2)
	// fmt.Println(a)
 
	fb:=fibo()
	for i:=0;i<10;i++{
		t:=fb()
		fmt.Println(t)
	}

}