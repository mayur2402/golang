package main;
import "fmt";

func factorial(num int) int{

	if(num > 1){
		return num * factorial(num - 1);
	}
	return num;
}

func main(){
	var num int = 0;
	var fact int = 0;

	fmt.Println("Enter number");
	fmt.Scanln(&num);

	if(num < 0){
		return;
	}

	fact = factorial(num);

	fmt.Println(fact);
}