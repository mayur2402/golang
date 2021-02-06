package main
import (
	"fmt"
	"io/ioutil"
)

func write(name string,data string){
	var tobyte = []byte(data)

	ioutil.WriteFile(name,tobyte,0777)

}

func main(){
	var name string
	var data string
	fmt.Println("Enter file name")
	fmt.Scanln(&name)

	fmt.Println("Enter data to write in file")
	fmt.Scanln(&data)

	fmt.Println(data)
	write(name,data)
}