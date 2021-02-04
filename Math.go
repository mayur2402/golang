package main
import (
	"fmt"
	"math"
)

func arithmatic(a int,b int){
	fmt.Printf("%d+%d = %d\n",a,b,a+b)
	fmt.Printf("%d-%d = %d\n",a,b,a-b)
	fmt.Printf("%d*%d = %d\n",a,b,a*b)
	fmt.Printf("%d/%d = %d\n",a,b,a/b)
}

func power(num float64,pow float64){

	ans := math.Pow(num,pow)

	fmt.Println(ans)
}

func squareroot(num float64){
	ans := math.Sqrt(num)

	fmt.Println(ans)
}

func main(){

	var a,b int
	var num,pow float64

	fmt.Println("Enter first number")
	fmt.Scanln(&a)

	fmt.Println("Enter second number")
	fmt.Scanln(&b)

	arithmatic(a,b)

	fmt.Println("Enter number to calculate power");
	fmt.Scanln(&num)
	fmt.Println("Enter power")
	fmt.Scanln(&pow)

	power(num,pow)

	fmt.Println("Enter number to calculate square root")
	fmt.Scanln(&num)

	squareroot(num)
}