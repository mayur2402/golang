package main
import "fmt"

func main(){

	for i := 0;i < 5;i++{
		fmt.Println(i);
	}

	for i := 0;i < 3;i++{
		for j := 3;j > 0;j--{
			fmt.Print(i,"\t",j);
		}
		fmt.Println()
	}

	for i:= 0; ;i++{
		fmt.Println(i);
		if(i == 50){
			break;
		}
	}

	i := 50
	for i > 0{
		fmt.Println(i);
		i--;
	}

	for true{
		fmt.Println("mayur")
	}
}