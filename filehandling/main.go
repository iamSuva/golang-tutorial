package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello world with go language")
	content := "Lets learn how to write on a file"

	file, err := os.Create("myfile.txt")
	if err != nil {
       panic(err)
	}
	len,err :=io.WriteString(file,content)
	if err !=nil{
		panic(err)
	}
	fmt.Println("length  of the file is ",len)
	defer file.Close()
	readFile ("myfile.txt")

}

func readFile(filename string){
databyte,err := ioutil.ReadFile(filename);
   if err !=nil{
	panic(err) //if error then panic will shut down execution
   }
   fmt.Println("file data is \n",string(databyte))
}
