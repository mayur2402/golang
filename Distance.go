/*
Write another program that converts from feet
into meters. ( 1 ft = 0.3048 m )
*/
package main;
import "fmt";

func main(){
	var feet float64;
	fmt.Println("Enter distance in feet");
	fmt.Scanln(&feet);

	var meter float64;

	meter = feet * 0.3048;

	fmt.Println("Distance in meter is ",meter);
}