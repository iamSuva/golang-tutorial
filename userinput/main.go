package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome:="welcome to user input";
    fmt.Println(welcome)

	reader :=bufio.NewReader(os.Stdin);
	fmt.Println("Enter the your age ");
	//comma ok | error ok

	input,_:=reader.ReadString('\n'); //read still \n
	fmt.Println("age is : ",input);
	fmt.Printf("type of age is %T\n ",input)
}		