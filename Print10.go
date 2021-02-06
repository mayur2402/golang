package main;
import "fmt";

func main(){
	i := 1;

	fmt.Println("Using for i <= 10");
	for i <= 10 {
		fmt.Println(i);
		i++;
	}

	fmt.Println("Using for i := 1; i <= 10; i++");
	for i := 1; i <= 10; i++ {
		fmt.Println(i);		
	}

	fmt.Println("Using for true");
	for true {
		fmt.Println(i);

		if(i == 1){
			break;
		}
		i--;
	}
}