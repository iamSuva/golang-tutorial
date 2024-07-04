package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web services")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() //closing the connection after getting response
  // to free the resources
	fmt.Printf("type of response %T\n", res)

	data,err:=ioutil.ReadAll(res.Body) // array of bytes
	
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(data))
}
