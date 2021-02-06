package main;
import (
	"fmt"
)

func main(){

	loop :
		fmt.Println("You are not eligible to vote");
		fmt.Println("Enter your age");
		var age int = 0;
		fmt.Scanln(&age);

		if(age <= 17){
			goto loop;
		}else{
			fmt.Println("You can vote");
		}
}