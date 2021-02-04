package main
import "fmt"

func main(){
	var num int = 0;
	fmt.Println("Enter Number");
	fmt.Scanln(&num);
	fmt.Println("You entered ",num);

	fmt.Println(num,"> 20",num > 20)
	fmt.Println(num,"< 20",num < 20)
	fmt.Println(num,"== 20",num == 20)
	fmt.Println(num,"!= 20",num != 20)
	
	var name string ;
	fmt.Printf("Enter name\t");
	fmt.Scanf("%s",&name);
	fmt.Printf("You entered %s\n",name);

	var age bool
	fmt.Print("Enter true if Your age is greater than 21 else false\n")
	fmt.Scan(&age);
	fmt.Print("Your age is greater than 21 ",age,"\n");
}