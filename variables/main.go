package main

import "fmt"
// name:="suvadip"  can't use outside method
const Token string="rtg3uhgigb3iugh"; //this is public as first letter is capital

func main(){
	// var username string ="suvadip";
	// fmt.Println(username);
	// fmt.Printf(" variable is type of %T\n",username);

	// var  isbool bool=true;
	// fmt.Println(isbool);
	// fmt.Printf("variable is of type %T\n",isbool);

	// var n uint8=255; // 8bits
	// fmt.Println(n);
	// fmt.Printf("variable is of type %T\n",n);
	// var num int=4555;
	// fmt.Println(num);
	// fmt.Printf("variable is of type %T\n",num);
	// var f float32=1523.5458216554455;
	// fmt.Println(f);
	// fmt.Printf("variable is of type %T\n",f);
	// var bigf float64=1523.5458216554455;
	// fmt.Println(bigf);
	// fmt.Printf("variable is of type %T\n",bigf);


	//implicit declaration
	var mynum int;
	fmt.Println(mynum); // default value is 0
	fmt.Printf("variable is of type %T\n",mynum);
    var myname string;
	fmt.Println(myname); // empty string
	fmt.Printf("varibale is of type %T\n",myname);
	//no variable styles

	user:=400;      // integer  // short declaration operator is used to declare the varible inside the function
	fmt.Println(user);
	fmt.Printf("variable is of type %T\n",user);
   

	fmt.Println(Token);
	fmt.Printf("variable is of type %T\n",Token);
	

}