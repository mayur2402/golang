package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	var inputReader *bufio.Reader

	inputReader = bufio.NewReader(os.Stdin)

//	fmt.Println(inputReader)

	fmt.Println("Enter full name")
	fname,err := inputReader.ReadString('\n')

	if err == nil {
		fmt.Println("Full name is ",fname)
	}

	var name = strings.TrimSpace(fname)

	fmt.Println(name)

/*
	var fullname string

	fmt.Println("Enter full name")
	fmt.Scanln(&fullname)

	fmt.Println(fullname)
	*/

	switch name {
	case "mayur" : fmt.Println("Hello "+name)
	}
}