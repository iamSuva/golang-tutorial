package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	fmt.Println("learning json")
	person:=Person{"suva",21,"male"}
	fmt.Printf("Person is %+v\n",person)
	//convert Person obj to json encoding (Marshal)
	jsondata,err:=json.Marshal(person)
	if err!=nil{
		fmt.Println("error marshalling json",err)
	}
	fmt.Println("person data is: ",string(jsondata))
    
	//decoding json data (unmarshalling)
	//convert json data back to person object
	var persondata Person
	errn:=json.Unmarshal(jsondata,&persondata)
	if errn!=nil{
		fmt.Println("error unmarshalling json",err)
	}

  fmt.Printf("person data after unmarshalling: %+v\n",persondata)
}
