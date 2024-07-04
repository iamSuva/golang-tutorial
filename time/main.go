package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println("current time is: ",currTime)
	formatedTime:=currTime.Format("02-01-2006,Monday,15:04:03 ,3:04 ")
	fmt.Println("formated time is: ",formatedTime)
    
 //my given date will be converted
	layout_str:="02-01-2006"
	mydate:="02-07-2024"
	formatedate,_ :=time.Parse(layout_str,mydate)
	fmt.Println("converted date is: ",formatedate)

}
