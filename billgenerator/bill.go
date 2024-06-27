package main

import (
	"fmt"
	"os"
)

type Bill struct {
	Name  string
	items map[string]float64
	tip   float64
}

// make a new bill
func hello() {
	fmt.Println("hello world")
}
func newBill(name string) Bill {
	bill := Bill{
		Name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return bill
}

//format the bill

func (b *Bill) formatbill() string{
	fs := "Bill breakdown: \n"
	var total float64=0

	//list items
	for k,v :=range b.items{
		fs+=fmt.Sprintf("%-20v........: $%v \n",k+" : ",v);
		total+=v;
	}
	total+=b.tip
	fs+=fmt.Sprintf("......the tip is:$%v\n",b.tip)
	fs+=fmt.Sprintf("total is .... : $%v  \n",total);
    return fs;
}

//update tip
func (b *Bill) updateTip(tip float64){
	b.tip=tip //struct pointers are auto,atically updated
}

//add item to items
func (b Bill) addItem(name string,price float64){
	b.items[name]=price;
}
func (b *Bill) save(){
	data:=[]byte(b.formatbill())
    err:=os.WriteFile("bills/"+b.Name+".txt",data,0644)
	if err!=nil{
		panic(err)
	}
	fmt.Println("saved")
	
}

