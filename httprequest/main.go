package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const URL="https://jsonplaceholder.typicode.com/posts/1"
func main() {
	fmt.Println("this is http request")
    response,err := http.Get(URL)
	if err!=nil{
        fmt.Println(err)
    
    }
	fmt.Printf("Response is type of : %T\n",response);
    result,err :=ioutil.ReadAll(response.Body)
	if err !=nil{
		panic(err)
	}
	content :=string(result);
	fmt.Println("content of get request is :\n",content)
	response.Body.Close()

}