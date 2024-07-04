package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello world with go language")
	reader := bufio.NewReader(os.Stdin);
	println("Enter a number to add with 5 : ");
	input,_ :=reader.ReadString('\n');
	println("Your entered number is : ",input);
    // num,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
	// if err!=nil{
	// 	fmt.Println(err);
	// } else
	// {
	// 	fmt.Println(num+5);
	// }
    num,_:=strconv.Atoi(strings.TrimSpace(input));
	fmt.Println("The sum is : ",num+5);
}