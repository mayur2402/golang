package main;
import "fmt";
func main(){
	var no1,no2,java int

	fmt.Printf("Enter first number")
	fmt.Scanf("%d",&no1)
	fmt.Printf("Enter second number")
	fmt.Scanf("%d",&no2)

	if (no1 == no2){
		fmt.Printf("%v and %v are equal\n",no1,no2);
	}else{
		fmt.Printf("%v and %v are not equal\n",no1,no2);
	}

	fmt.Printf("Enter rating out of 5 for java")
	fmt.Scanf("%d",&java);

	switch java{
	case 0:
		fmt.Printf("Bad\n")
	case 1:
		fmt.Printf("beginner\n")
	case 2:
		fmt.Printf("Average\n")
	case 3:
		fmt.Printf("intermediate\n")
	case 4:
		fmt.Printf("Perfect\n")
	case 5:
		fmt.Printf("Expert\n")
	default:
		fmt.Printf("Invalid rate\n")
	}
	
}