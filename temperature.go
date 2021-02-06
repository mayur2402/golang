/*
Using the example program as a starting point,
write a program that converts from Fahrenheit
into Celsius. ( C = (F - 32) * 5/9 )
*/

package main;
import "fmt";

func main(){
	var farh float64;
	fmt.Println("Enter temperature in Fahrenheit");
	fmt.Scanln(&farh);

	var cel float64;

	cel = ((farh - 32) * 5) / 9;

	fmt.Println("Temperature in Celsius is ",cel);
}