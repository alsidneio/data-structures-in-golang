package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Append(num int) {
	node := &Node{value: num}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{}
	for i := 0; i < 10; i++ {
		value := randInt(-100, 100)

		fmt.Printf("Adding the value %d to the list \n", value)
		l.Append(value)
	}

}

// helper functions
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func insertNode(current *ListNode, new *ListNode){
    next := current.Next
    current.Next = new
    new.Next = next
}

func TraversList(l *List, val int ){
	
	// fix this fuction so that we accept a value we are looking for. 
	
	for {

		// Checking for an empty list
		if l.head == nil {
			fmt.Println("list is empty")
			break
		}

		fmt.Printf("Current node value is %d\n", l.head.value)

		//Checking to see if we are at the end of the list
		if l.head.next == nil {
			break
		}

		// Move to the next node
		l.head = l.head.next
	}
}