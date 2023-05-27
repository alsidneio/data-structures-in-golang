package main

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	prev := new(ListNode)

	// may need to traverse until you get to a non duplicate
	for curr != nil && curr.Next != nil && curr.Next.Val == curr.Val {
		curr.Next = curr.Next.Next
	}

	// loop for happy path & nil
	for curr != nil {
		if curr.Val == prev.Val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return head

}
