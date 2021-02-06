package main;
import "fmt";

func swapByAddress(n *int,n2 *int){
	var temp int = 0;
	temp = *n;
	*n = *n2;
	*n2 = temp;
}

func swapByValue(n int,n2 int){
	var temp int = 0;
	temp = n;
	n = n2;
	n2 = temp;
}

func main(){
	var num int = 0;
	var num2 int = 0;

	fmt.Println("Enter first number");
	fmt.Scanln(&num);

	fmt.Println("Enter second number");
	fmt.Scanln(&num2);

	fmt.Println("Swap by address");

	fmt.Println(num," ",num2);

	swapByAddress(&num,&num2);

	fmt.Println(num," ",num2);

	fmt.Println("Swap by value");

	fmt.Println(num," ",num2);

	swapByValue(num,num2);

	fmt.Println(num," ",num2);

	fmt.Println(&num," ",&num2);
}