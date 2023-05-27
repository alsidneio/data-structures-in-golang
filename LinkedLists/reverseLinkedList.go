package main

/**

Strategy
---
use 3 pointers:
1. curr: Node currently being reversed
2. next: Node that is saved to preserve chain of nodes to be worked on next
3. previous: node that the current node points to in reversion. It also saves the reversed node history.

Basic Idea
---
As we traverse down the list. We set the ListNode.Next pointer to the previous node.
we know we have reached the end when have reached the end of the list when current node is equal to nil

Psuedo Code
---
check for nil head
	if nil return head

for each iteration:
	save the downstream node
	set the current.next point to previous node
	move previous node to current node
	move head to the current node
	move current node to the next node

	if current == nil
		return the head

**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	var current = head
	var next *ListNode

	if current == nil {
		return head
	}

	// loop to continually reverse
	for {

		next = current.Next
		current.Next = previous
		previous = current
		head = current
		current = next

		if current == nil {
			return head
		}
	}

}
