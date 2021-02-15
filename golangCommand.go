package main

import (
	"fmt"
	"os"
)

func main() {

	i := 1
	for i < len(os.Args) {
		fmt.Println(os.Args[i])
		i++
	}

	fmt.Println("length is ",len(os.Args))
}
