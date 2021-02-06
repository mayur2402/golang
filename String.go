package main
import (
	"fmt"
	"strconv"
	"strings"
)

func conactenation(){
	var str1 string = "Mayur"
	var str2 string = "Dimble"

	fmt.Println(str1+str2)

	str := fmt.Sprintf("%s %s",str1,str2)
	fmt.Println(str)
}

func stringToInt(){
	var num = "20"
	var no int64
	var err error

	no,err = strconv.ParseInt(num,10,64)

	if err == nil{
		fmt.Println(no)
	}else{
		fmt.Println(err)
	}
}

func stringToFloat(){
	var num = "20.20"
	var no float64
	var err error

	no,err = strconv.ParseFloat(num,64)

	if err == nil{
		fmt.Println(no)
	}else{
		fmt.Println(err)
	}
}

func numberToFloat(){
	var num = 40
	var str string
	
	str = fmt.Sprintf("%d",num)

	fmt.Println(str)
}

func copydata(){
	var str = "hello world"
	var cstr = str[:len(str)]
	fmt.Println(cstr)
	cstr = str[0:5]
	fmt.Println(cstr)
	cstr = str[6:len(str)]
	fmt.Println(cstr)
}

func main(){
	var str string = "hello world"
	fmt.Println(len(str))
	fmt.Println(str)
	fmt.Println(strings.ToUpper(str))

	str = "HOW ARE YOU"
	fmt.Println(str)
	fmt.Println(strings.ToLower(str))

	conactenation()
	stringToInt()
	stringToFloat()
	numberToFloat()
	copydata()
}