package main

import (
	"fmt"
	// "sort"
)
func main() {
	fmt.Println("Hello world with go language")
	// var fruitlist=[]string{"apple","gouva"};
	// fmt.Println(fruitlist);
	// fruitlist=append(fruitlist,"mango","lichi");
	// fmt.Println(fruitlist);

	// // fruitlist=append(fruitlist[1:] );
	// fruitlist=append(fruitlist[:3] );
	// fmt.Println(fruitlist);
	
	// marks:=make([]int,4)
	// marks[0]=70
	// marks[1]=65
	// marks[2]=85
	// marks[3]=75
	// fmt.Println(marks);
	// marks = append(marks, 95,88,63);
	// fmt.Println(marks);
	// sort.Ints(marks);
	// fmt.Println(marks);

	var course=[]string{"html","js","react","next","nodejs"};
	fmt.Println(course)
	var ind int=3;
	course=append(course[:ind],course[ind+1:]... )
	fmt.Println(course)
}
