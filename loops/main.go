package main

import "fmt"

func main() {
	fmt.Println("Hello world with go language")
	//arr:=[]int{10,20,30,40,50};
	// for i:=range arr{
	// 	fmt.Println(arr[i]);
	// }

	// for _,val:=range arr{
	// 	fmt.Println(val);
	// }
	//j:=0;
	// for j<5{
	// 	if j==3{

	// 		break;
	// 	}
		
	// 		println(j,"-> ",arr[j]);
	// 		j++;
		
	// }
//   for i:=0;i<=10;i++{
// 	fmt.Print(i," ")
//   }
  myarr:=[]int{10,20,30,40,50}
  for ind,val:=range myarr{
	fmt.Println("ind and  value is: ",ind,val);
  }
}
