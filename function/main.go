package main

import "fmt"

func main() {
	fmt.Println("Hello this is function")
	// greet()
	ans,name:=sum(10,15)
	fmt.Println("sum is ",ans ,"and name is ",name);
	res:=allsum(1,2,3,5,8)
	fmt.Println("sum is ",res);
}
func sum(a int, b int)(int,string){
	return a+b,"suvadip";
}
func allsum(arr ...int)int{
	res:=0
	for _,val:=range arr{
		res+=val;
	}
	return res
}
func greet(){
	fmt.Println("welcome to function");
}