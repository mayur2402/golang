package main
import (
	"fmt"
	"time"
)

type Employee struct{
	eid int
	ename string
	edate time.Time
}

func main(){
	var num int = 0
	fmt.Println("How many employee")
	fmt.Scanln(&num)

	for i := 0;i < num;i++ {
		emp := new(Employee)
		fmt.Println("Enter id")
		fmt.Scanln(&emp.eid)
		fmt.Println("Enter name")
		fmt.Scanln(&emp.ename)
		emp.edate = time.Now()

		fmt.Println(emp.eid,emp.ename,emp.edate)
	}
}