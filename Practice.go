package main

import "fmt"

func main(){
	var num int = 20

	fmt.Println(num)

	fmt.Printf("Variable num is of type %T\n",num)

	num1 := 40.40

	fmt.Println(num1)

	fmt.Printf("Variable num1 is of type %T\n",num1)

	var num2 float32 = 40.40

	fmt.Printf("%f\n",num2)

	var num3 float64 = 40.40

	fmt.Printf("%f\n",num3)

	fmt.Println(num3)

	const PI float32 = 3.142

	fmt.Println(PI)

	num++

	fmt.Println("num++",num)

/*	PI++

	fmt.Println("PI++",PI)
	*/
}