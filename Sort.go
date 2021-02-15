package main
import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	length := len(os.Args) - 1
	if length == 0 {
		fmt.Println("Enter numbers to sort")
		return
	}

	var i int = 1
	var j int = 0
	arr := make([]int64,length)
	//var arr []int64
	for i < len(os.Args) {
		arr[j],_ = strconv.ParseInt(os.Args[i],10,64)
		i++
		j++
	}
	j = 0
	i = 0
	

	for i < length {
		for j = 0;j < length-1-i;j++ {
			
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		i++
	}
	
	
	fmt.Println(arr)
}