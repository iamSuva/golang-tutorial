package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"os"
)

func main() {
// 	file, err := os.Create("myfile.txt")
// 	if err!=nil{
// 		fmt.Println("error while creating file")
// 		return;
// 	}
// 	defer file.Close()
// 	fmt.Println("file is created")
// 	content:="this is myfile.txt"
//    byte,errs:=io.WriteString(file,content)
//    fmt.Println(byte)
//    if errs!=nil{
//     fmt.Println("error while writing content to file")
//    }
//    fmt.Println("content written successfully")

  readmyfile("myfile.txt")

}
func readmyfile(s string){
 file,err:=os.Open(s);
 if err!=nil{
	 fmt.Println("error while opening file")
	 return
 }
 defer file.Close()
 //create buffer to read file content
 buffer:=make([]byte,1024);
 //read the file content
 for{
	n,err:=file.Read(buffer)
	if err==io.EOF{
        break;
    }  
	if err!=nil{
		fmt.Println("error reading")
		break;
	}
	fmt.Println(string(buffer[:n]))
 }
}

func readFile(s string){
	// conent,err:=ioutil.ReadFile(s); // it is not good for large file
	conent,err:=os.ReadFile(s); // it is not good for large file
	if err!=nil{
		fmt.Println("error while reading file")
	}
	fmt.Println(string(conent))

}