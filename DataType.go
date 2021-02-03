package main;
import "fmt";

func main(){
	var no int8 = 20;

	fmt.Println(no);

	var no1 int16 = 200;

	fmt.Println(no1);

	var no3 float32 = 200.2002;

	fmt.Println(no3);

	var no4 float64 = 200.2002002;

	fmt.Println(no4);

	var sum = no1 + 20;

	fmt.Println("200 + 20 = ",sum);

	var sub = no - 10;

	fmt.Println("20 - 10 = ",sub);

	var div = no3 / 20;

	fmt.Println("200.20002 / 20 = ",div);

	var str string = "hello";

	fmt.Println(str);

	var str2 string = "world";

	fmt.Println(str2);

	var s = str+" "+str2;

	fmt.Println(s);

	fmt.Println("len(hello)",len(str));

	fmt.Println("len(world)",len(str2));

	fmt.Println("len(hello world)",len(s));

	fmt.Println("hello[2]",str[2]);
}