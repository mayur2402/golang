package main;
import "fmt";

func main(){
	var lang = map[string]map[string]string{
		"aniket" :  map[string]string{
			"mid level" : "C",
			"oo" : "C++",
			"vm" : "java",
		},
		"Devarshi" : map[string]string{
			"oo" : "python",
			"script" : "javascript",
		},
		"sagar" : map[string]string{
			"script" : "php",
			"new" : "scala",
		},
	}

	fmt.Println(lang);
}