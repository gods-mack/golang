
package main

import ("fmt") 


type node struct {
	data int
	next *node
}


var head *node = nil
var tail *node = nil


func insert_node(x int) {
	
	if head == nil {
		head =  &node{data:x, next:nil}
		tail = head
	} else {
		tail.next = &node{data:x, next:nil}
		tail = tail.next
	}
}

func main() {
	//link := List{}
	insert_node(32)
	insert_node(53)
	insert_node(12)
	insert_node(19)
	insert_node(78)
	insert_node(10)
	

	// traverse linked list
	h := head
	for h!=nil {
		fmt.Println(h.data, " ")
		h = h.next
	}
}