package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func createBill() Bill{
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("create a new bill for ")
	name,_:=reader.ReadString('\n')
	name=strings.TrimSpace(name)
	b:=newBill(name)
	fmt.Println("bill is created for : ",b.Name)
	return b;

}
func getOptions(b Bill){
	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("Choose option ( a) -> add item , (s) save bill , (t) -add tip")
	opt,_:=reader.ReadString('\n')
	opt=strings.TrimSpace(opt)
	switch(opt){
	case "a":
		fmt.Println("add item to bill")
		fmt.Println("enter the item name :")
		name,_:=reader.ReadString('\n')
		fmt.Println("enter the item price :")
		price,_:=reader.ReadString('\n')
		//fmt.Println("type of name and price is %T %T",name,price)
		p,err:=strconv.ParseFloat(strings.TrimSpace(price),64);
		if err!=nil{
			// log.Fatal(err)
			fmt.Println("price must be a number")
			getOptions(b)
		}
		b.addItem(name,p)
		fmt.Println("item is added ",name,price);
		getOptions(b);
	case "t":
		fmt.Println("enter the tip")
		tip,_:=reader.ReadString('\n')
		t,err:=strconv.ParseFloat(strings.TrimSpace(tip),64)
		if err!=nil{
            // log.Fatal(err)
            fmt.Println("tip must be a number")
            getOptions(b)
        }
		b.updateTip(t)
		fmt.Println("tip is added ",tip);
		getOptions(b);
	case "s":
		b.save();
		
		fmt.Println("save bill ",b.Name)
    default:
		fmt.Println("invalid option")
		getOptions(b)						
	}
//    fmt.Println(opt)
}
func main() {
    fmt.Println("Hello world with go language")
    // myBill := newBill("suvadip")
    // myBill.updateTip(20)
	mybill:= createBill();
	getOptions(mybill)
	
	
}