package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	fmt.Println("Hello world with go language")
	// rand.Seed(time.Now().UnixNano());
	dicenum:=rand.Intn(6)+1;
	println(dicenum);
	switch dicenum{
	case 1:
		fmt.Println("You got 1");
	case 2:
		fmt.Println("You got 2");
	case 3:
		fmt.Println("You got 3");
	case 4:
		fmt.Println("You got 4");
	case 5:
		fmt.Println("You got 5");
	case 6:
		fmt.Println("You got 6");
    default :
	  fmt.Println("game over")		


	}
}
