package main
import( 
	"fmt"
)

type rectangle struct{
	x,y float64
}

func area(r rectangle) float64 {
	return r.x*r.y
}

func (r rectangle) area() float64{
	return r.x*r.y
}

func main(){
	rect := rectangle{5,3}
	a := area(rect)
	fmt.Println("Area is",a)

	b := rect.area()
	fmt.Println("Area is",b)
}