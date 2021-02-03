package main;
import "fmt";

func main(){
	var arr[5] string;

	arr[4] = "golang";

	fmt.Println(arr);

	var brr[5] int;

	brr[0] = 20;
	brr[1] = 30;
	brr[2] = 40;
	brr[3] = 50;
	brr[4] = 60;

	fmt.Println(brr);

	var ans int = 0;
	for i := 0;i < 5;i++ {
		ans = ans + brr[i];
	}

	fmt.Println("Sum of array elements is",ans);

	crr := [5]float64{98,78.987,956.86,435.89,23.9};

	fmt.Println(crr);

	var total float64 = 0;
	for _,value := range crr {
		total = total + value;
	}

	fmt.Println("Sum of array elements is",total);

	fmt.Println("crr[2:]",crr[2:]);

	fmt.Println("crr[:3]",crr[:3]);


	err := []float64{200,300,400,500,600};

	fmt.Println(err);

	frr := err[2:4];

	fmt.Println("err[2:4]",frr);


	drr := make([]float64,6)

	drr[0] = 10;

	fmt.Println(drr);


	grr := append(err,7,9);

	fmt.Println("grr := append(err,7,9)",grr);

	copy(drr,err);

	fmt.Println(drr);


	var x map[string]int;

	fmt.Println(x);
	

	y := make(map[string]int);

	y["one"] = 1;

	fmt.Println(y["one"]);


	i := make(map[int]int);

	i[20] = 20;

	fmt.Println(i[20]);

	fmt.Println(i);
	
	for j := 2; j < 5;j++ {
		i[j] = j;
	}

	fmt.Println(i);

	name , k := i[2];

	fmt.Println(name,k);


	var Map = map[string]string{
		"aniket" : "java",
		"devarshi" : "cpp",
		"sagar" : "c",
	}

	fmt.Println(Map);
}