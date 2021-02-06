package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func rating(input *bufio.Reader,err error){

	var fullname string
	fmt.Println("Enter your full name")
	fullname,err = input.ReadString('\n')

	if err == nil {

		name := strings.TrimSpace(fullname)

		fmt.Println("Hello "+name)
		fmt.Println("Please rate our pizza out of 5")

		rating,err := input.ReadString('\n')

		if err == nil {

			rate,err := strconv.ParseInt(strings.TrimSpace(rating),10,64)

			if err == nil {
				displayMessage(name,rate)
			}
		}
	}
}

func displayMessage(fullname string,rate int64){

	if rate == 5{
		fmt.Println(fullname+" our staff will get bonus for that")
	} else if rate == 4 || rate == 3 {
		fmt.Println(fullname+" we are working diligently to be the best")
	} else if rate == 2 || rate == 1 || rate == 0{
		fmt.Println(fullname+" we will try to be better")
	} else {
		fmt.Println("Thank you !")
	}
}

func main(){
	
	var inputReader *bufio.Reader
	var err error

	inputReader = bufio.NewReader(os.Stdin)

	rating(inputReader,err)

}