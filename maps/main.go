package main

import "fmt"

func main() {
	fmt.Println("Hello world with go language")
// 	var languanges = make(map[string]string)
// 	languanges["js"] = "javascript"
// 	languanges["py"] = "python"
// 	languanges["rb"] = "ruby"
// 	languanges["cpp"] = "C++"

// 	fmt.Println("map is ",languanges);
// delete(languanges,"rb");
// for key,val:=range languanges{
// 	fmt.Printf("key %v and value %v is \n",key,val)
// }
  students :=map[int]string{
  1:"sachin",
  2:"rahul",
  3:"virat",
  }
  fmt.Println(students);
  students[1]="suvadip"
  fmt.Println(students);

}