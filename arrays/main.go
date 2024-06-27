package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang");
	// var fruits [4]string;
	// fruits[0]="Apple"
	// fruits[1]="Mango"
	// fruits[3]="ornage"
	// fmt.Println("fruit list is ",fruits);
	// fmt.Println("fruit list is ",len(fruits));

	// var veg=[3]string{"potato","beans","tomato"};
	// fmt.Println("veg list is ",veg);
//    var arr=[3]int{10,20,30};
//    fmt.Println(arr);
//    friends:=[2]string{"suva","jeet"};
//    fmt.Println(friends,len(friends));


   //slices  dont have fixed length
//    var marks []int;
//    //append method return a new array
//    marks=append(marks,10,20,30,40,50);
//    fmt.Println(marks);
//   marks[0]=100
//   fmt.Println(marks);

  var score=[]int{5,8,9,6}
  fmt.Println(score)
  score=append(score, 7)
  fmt.Println(score)
  range1:=score[1:3]
  range2:=score[1:]
  fmt.Println(range1,range2)
}