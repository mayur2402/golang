package main
import(
	"fmt"
	"math"
)

func area(radius float64){
	area := math.Pi*math.Pow(radius,2)

	fmt.Printf("Area of circle is %.3f\n",area)
}

func main(){
	var rad float64;

	fmt.Println("Enter radius of circle")
	fmt.Scanln(&rad)

	area(rad)
}