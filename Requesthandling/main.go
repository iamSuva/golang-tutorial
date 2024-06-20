package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to request handling")
	//getRequest()
	// postRequest()
	postFormdata()
}

func getRequest() {
	const url string = "http://localhost:4000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(result)
	fmt.Println(content)
}


func postRequest(){
	const url string ="http://localhost:4000/post"
	requestBody :=strings.NewReader(`
	{
	"coursename":"mycourse",
	"price":0,
	"platform":"learnme"
	}
	`)

	response,err:=http.Post(url,"application/json",requestBody);
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func postFormdata(){
	const myurl string="http://localhost:4000/postform"

	data := url.Values{}
	data.Add("name","suva")
	data.Add("age","21")
	data.Add("email","suva@gmail.com")
    
	response,err:=http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))

}