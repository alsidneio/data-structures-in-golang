package main

// code needs to be optimized, alot of repeating statments
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head

	// check for empty conditions
	if list1 == nil && list2 == nil {
		return head.Next
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	// traverse through both lists
	for {
		if (list1.Val < list2.Val) || (list1.Val == list2.Val) {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next

		if list1 == nil && list2 == nil {
			return head.Next
		}

		if list1 == nil {
			curr.Next = list2
			return head.Next
		}

		if list2 == nil {
			curr.Next = list1
			return head.Next

		}
	}
}
