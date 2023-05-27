package main
/**
 *   This is just an advanced delete opertation
 *   the only gotcha in this is if the deletion node is at the head
 *   in that case, one will need to traverse down the node until we get
 *   a non-deletion value then start the nil and happy path pattern. 
 **/





 func removeElements(head *ListNode, val int) *ListNode {
    curr := head 
    var next *ListNode
    prev := new(ListNode)

for  head != nil && head.Val == val{
    head = head.Next
}

// this for loop handles a nil and happy path conditions 
    for curr != nil{
        if  curr.Val != val{
            prev = curr
            curr = curr.Next
        }else{   
            next = curr.Next
            prev.Next = next
            curr = next
        }
    }

    return head 
    
}