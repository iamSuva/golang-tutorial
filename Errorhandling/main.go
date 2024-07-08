package main

import (
	"errors"
	"fmt"
)

func main() {
	// str := "suva"
	// err := check(str)
	// if err != nil {
	// 	fmt.Println("not matched")
	// }else{
	// 	fmt.Println("matched")
	// }
	res,err:=division(25,10);
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(res)
	}
}
func check(s string) error {
    myerr:=errors.New("invalid string")
	if s!="suvad"{
		return myerr
	}
	return nil
}
func division(a, b int) (int,error){
      if b==0{
		return 0,fmt.Errorf("denominator can not be %d",b)
	  }else{
		return a/b,nil
	  }
	  
}