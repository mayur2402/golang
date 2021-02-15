package main;

import (
	"fmt"
)

type node struct{
	data int
	next *node
}

func insertFirst(first **node,num int){
	fmt.Println(*first,num)
	if *first == nil {
		*first = new(node)
		(*first).data = num
		(*first).next = nil
	} else {
		temp := *first
		
		*first = new(node)
		(*first).data = num
		(*first).next = temp
	}
}

func insertLast(first **node,num int) *node{

	var temp **node = first
	var add *node = *first
	
	if *temp == nil {
		fmt.Println(first,num)
		insertFirst(first,num)
	} else {
		
		for (*temp).next != nil {
			*temp = (*temp).next
		}
		(*temp).next = new(node)
		(*temp).next.data = num
		(*temp).next.next = nil
	}	
	return add
}

func insertAtPos(first *node,pos int,num int){

	var temp *node
	size := count(first)
	if pos == 1 {
		insertFirst(&first,num)
	} else if pos == (size+1){
		insertLast(&first,num)
	} else if pos <= size {
		
		var i int = 1
		temp = first
		for i < pos-1 {
			temp = (temp).next
			i++
		}

		newn := new(node)
		newn.data = num
		newn.next = temp.next
		temp.next = newn
	} else {
		fmt.Println("Invalid position")
	}
}

func deleteFirst(first **node){

	if *first == nil {
		fmt.Println("invalid position")
		return
	}

	if (*first).next == nil {
		*first = nil
	} else{
		*first = (*first).next
	}
}

func deleteLast(first **node){

	if *first == nil {
		fmt.Println("invalid position")
		return
	}

	if (*first).next == nil {
		*first = nil
	} else {
		for (*first).next.next != nil {
			*first = (*first).next
		}
		(*first).next = nil
	}
}

func deleteAtPos(first *node,pos int){

	var size = count(first)

	if first == nil {
		fmt.Println("invalid position")
		return
	}

	var temp *node
	temp = first
	if pos > size {
		return
	} else if pos == 1 {
		deleteFirst(&first)
	} else if(pos == size) {
		deleteLast(&first)
	} else {
		var i int = 1
		for i < pos-1 {
			temp = temp.next
			i++
		}
		temp.next = temp.next.next
	}
}

func display(first *node){

	fmt.Print("Elements in the list\t = ")
	for first != nil {
		fmt.Printf("%v\t",first.data)
		first = first.next
	}
	fmt.Println()
}

func count (first *node) int{
	var count int = 0

	for first != nil {
		count++
		first = first.next
	}
	return count
}

func main(){
	var head *node = nil
	var num int = 0
	var ch int = 0

	for ch != 9 {
		fmt.Println("----------------------------------------------")
		fmt.Println("1.insertFirst\n2.insertLast\n3.insertAtPos\n4.deleteFirst\n5.deleteLast")
		fmt.Println("6.deleteAtPos\n7.display\n8.count\n9.exit")
		fmt.Println("----------------------------------------------")
		fmt.Println("Enter your choice")
		fmt.Scanln(&ch)

		switch ch {
		case 1 : 
			fmt.Println("Enter number to insert first position")
			fmt.Scanln(&num)

			insertFirst(&head,num)

		case 2 :
			fmt.Println("Enter number to last position")
			fmt.Scanln(&num)

			head = insertLast(&head,num)

		case 3 :
			var pos,num int
			fmt.Println("Enter position")
			fmt.Scanln(&pos)
			fmt.Println("Enter number")
			fmt.Scanln(&num)

			insertAtPos(head,pos,num)

		case 4 :
			deleteFirst(&head)

		case 5 :
			deleteLast(&head)			

		case 6 :
			var pos int
			fmt.Println("Enter position")
			fmt.Scanln(&pos)
			
			deleteAtPos(head,pos)

		case 7 :
			display(head)

		case 8 :
			fmt.Println("Number of elements in list ",count(head))

		}
	}
}