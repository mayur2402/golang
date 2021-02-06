package main;
import "fmt";

func display(arr []int){
	for _,j := range arr{
		fmt.Printf("%d\t",j);
	}
	fmt.Println(); 
}

func average(arr []int) int {
	var add int = 0;

	for _,j := range arr {
		add = add + j;
	}

	return add/len(arr);
}

func math(a []int) (int,int,int) {

	var add int = 0;
	var sub int = 0;
	var mul int = 1;

	for i := 0;i < len(a);i++ {
		add = add + a[i];
		sub = sub - a[i];
		mul = mul * a[i];
	}
	return add,sub,mul;
}

func friend(list ...string){
	for _,name := range list{
		fmt.Println(name);
	}
}

func oddnumber() func() int{
	i := 1;

	return func() int{
		i = i + 2;
		return i;
	}
}
func main(){
	
	arr := []int{20,30,40,50};

	for i := 0;i < len(arr);i++ {
		fmt.Printf("%d\t",arr[i]);
	}
	fmt.Println();

	for _,j := range arr {
		fmt.Print(j,"\t");
	}
	fmt.Println();

	display(arr);

	var avg = average(arr);

	fmt.Println(avg);

	var add,sub,mul = math(arr);

	fmt.Println(add);

	fmt.Println(sub);

	fmt.Println(mul);

	friend("aniket","devarshi","sagar");

	var age int = 21;

	year := func(age int) int{
		age++;
		return age;
	}

	fmt.Println(year(age));

	even := func(num int) {
		for i := 2; i < num;i++ {
			if(i % 2 == 0){
				fmt.Printf("%d\t",i);
			}
		}
	}

	even(20);

	fmt.Println();

	o := oddnumber();

	fmt.Println(o());
	fmt.Println(o());
}