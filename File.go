package main
import(
	"fmt"
	"io/ioutil"
)

func writeFile(str string){
	bytearr := []byte(str)
	ioutil.WriteFile("m.txt",bytearr,0777)
	fmt.Println("File created")
}

func readFile(){
	data,_ := ioutil.ReadFile("m.txt")
	fmt.Println(string(data))
}

func main(){
	writeFile("Hello world")
	readFile()
}

