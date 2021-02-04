package main
import (
	"fmt"
)

func main(){
	var num int = 10

	fmt.Println("Value = ",num)
	fmt.Println("Address = ",&num)

	var n int = 20
	var ptr *int 
	ptr = &n

	fmt.Println("Value= ",*ptr)
	fmt.Println("Address = ",ptr)

	no := 30
	iptr := &no

	fmt.Println("Value = ",*iptr)
	fmt.Println("Address = ",iptr)

	var number int = 40
	var p = new(int)
	p = &number;

	fmt.Println("Value = ",*p)
	fmt.Println("Address = ",p)

}