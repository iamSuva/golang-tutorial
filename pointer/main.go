package main


func main(){
	println("Welcome to pointer");
	var num int=25;
	var ptr=&num;
	println("address of ptr is ",ptr);
	println("value of ptr is ",*ptr);

	*ptr=*ptr *4;
	println("value of ptr and num  is ",*ptr,num);
}