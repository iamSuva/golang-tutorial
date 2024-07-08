package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func postRequest() {
	myurl := "https://jsonplaceholder.typicode.com/todos"
	mytodo := Todo{
		UserId:    10,
		Title:     "new task mine",
		Completed: false,
	}
	//convert obj to json
	jsondata, err := json.Marshal(mytodo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}
	//convert json data to string
	jsonstring := string(jsondata)
	//convert string data to reader
	jsonreader := strings.NewReader(jsonstring)

	res, err := http.Post(myurl, "application/json", jsonreader)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("response status: ", res.Status)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))
}
func updateRequest() {
	todo := Todo{
		UserId:    5525,
		Title:     "updated task mine",
		Completed: true,
	}
	jsondata, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}
	jsonstring := string(jsondata)
	jsonreader := strings.NewReader(jsonstring)
    
	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodPut, myurl, jsonreader)
	if err != nil {
		fmt.Println("Error making request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res,err:=client.Do(req)
	if err!= nil{
		fmt.Println("Error making request: ", err)
    }
	defer res.Body.Close()
    
	fmt.Println("response status: ", res.Status)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))
}
func deleteRequest(){
	myurl := "https://jsonplaceholder.typicode.com/todos/1"
    req,err:=http.NewRequest(http.MethodDelete,myurl,nil);
	if err!= nil{
        fmt.Println("Error making request: ", err)
    }
	// req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res,err:=client.Do(req)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response data: ", string(data))
	fmt.Println("response: ", res.Status)
}



func main() {
	//postRequest()
//    updateRequest()
  deleteRequest()
}
