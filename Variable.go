package main

import "fmt"

func main(){

	var i int;

	var f float32;

	var str string;

	var b bool;

	fmt.Printf("%T\n%T\n%T\n%T\n",i,f,str,b);

	fmt.Printf("%v\n%v\n%v\n%v\n",i,f,str,b);
}