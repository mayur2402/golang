package main
import (
	"fmt"
	"os"
	"strconv"
	)
	
func main() {

	if len(os.Args) != 2 {
		fmt.Println("Enter only one argument")
		return
	}
	
	num,err := strconv.ParseInt(os.Args[1],10,64)
	
	if(err == nil) {
		var i int64
		for i = 1;i <= 10; i++ {
			fmt.Println(i * num)
		}
	} else {
		fmt.Println(err)
	}
}
