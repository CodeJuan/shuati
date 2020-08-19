package merge_two_sorted_lists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var l *ListNode
	var current *ListNode

	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			if l2 == nil {
				l, current = setReturnList(l1, l, current)
				l1 = l1.Next
			} else {
				if l1.Val <= l2.Val {
					l, current = setReturnList(l1, l, current)
					l1 = l1.Next
				} else {
					l, current = setReturnList(l2, l, current)
					l2 = l2.Next
				}
			}
		} else {
			l, current = setReturnList(l2, l, current)
			l2 = l2.Next
		}
	}
	return l
}

func setReturnList(inputList *ListNode, l *ListNode, current *ListNode) (*ListNode, *ListNode) {
	if l == nil {
		l = &ListNode{inputList.Val, nil}
		current = l
	} else {
		current.Next = &ListNode{inputList.Val, nil}
		current = current.Next
	}
	return l, current
}
