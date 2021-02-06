package main
import (
	"fmt"
	"math/rand"
)

type Node struct{
	data int
	next *Node
}

func add(first *Node,data int) *Node {
	if first == nil{
		first = new(Node)
		first.data = data
		first.next = nil
	}else{
		temp := new(Node)
		temp.data = data
		temp.next = first

		first = temp
	}

	return first
}

func display(first *Node){
	var temp *Node 
	temp = first

	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main(){

	var head *Node
	head = nil

	for i := 0;i < 5;i++ {
		head = add(head,rand.Intn(20000))
	}

	display(head)
}