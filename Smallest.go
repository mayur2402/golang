package main;
import "fmt";

func main(){
	list := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97, 9,17,}

	fmt.Println(list);

	var max int = 0;
	var min int = list[0];

	for i := 0; i < len(list); i++ {
		if(max < list[i]){
			max = list[i];
		}

		if(min > list[i]){
			min = list[i];
		}
	}

	fmt.Println(max);
	fmt.Println(min);
}