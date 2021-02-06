package main;
import "fmt";

func main(){

	fmt.Println("Array define")
	var arr[5] int 		//one dimentional array
	var arr2[3][3] int 	//two dimentional array

	fmt.Println("Slice define")
	var Arr[] int
	Arr2 := make([]int,5)
	Arr3 := make([][]int, 3*3)

	fmt.Println("map define")
	Map := make(map[string]int)
	var Map2 map[string]int

	fmt.Println(arr,arr2,Arr,Arr2,Arr3,Map,Map2)
}