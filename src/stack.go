package main

import ("fmt")

type node struct {
	data int
	next *node
}


type stack struct {
	top *node
}


func push(x int, stak *stack) {

	if stak.top == nil {
		stak.top= &node{data:x, next:nil}
	} else{
		stak.top =  &node{data:x, next:stak.top}

	}

}


func pop(stak *stack) {

	if stak.top == nil {
		fmt.Println("Stack is empty!")
		return
	} else {
		stak.top = stak.top.next
	}
}



func main() {

	mystack := stack{top:nil, head:nil} 
	push(10, &mystack)
	push(15, &mystack)
	push(16, &mystack)
	push(23, &mystack)
	pop(&mystack)
	push(76, &mystack)
	pop(&mystack)

	fmt.Println("stack top: ", mystack.top.data)

}
