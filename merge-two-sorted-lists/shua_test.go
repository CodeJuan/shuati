package merge_two_sorted_lists

import (
	"fmt"
	"testing"
)

type ucTest struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

var ucTests = []ucTest{
	{buildList([]int{2}),
		buildList([]int{1}),
		buildList([]int{1, 2})},
	{buildList([]int{1, 3}),
		buildList([]int{2, 4}),
		buildList([]int{1, 2, 3, 4})},
	{nil, nil, nil},
	{buildList([]int{1, 2}),
		nil,
		buildList([]int{1, 2})},
	{nil,
		buildList([]int{1, 2}),
		buildList([]int{1, 2})},
}

func buildList(in []int) *ListNode {
	var l *ListNode
	var cur *ListNode
	for _, v := range in {
		if l == nil {
			l = &ListNode{v, nil}
			cur = l
		} else {
			tmp := &ListNode{v, nil}
			cur.Next = tmp
			cur = tmp
		}
	}
	return l
}

func printList(l *ListNode) {
	fmt.Print("new list\n")
	if l == nil {
		fmt.Print("nil list\n")
	}
	for l != nil {
		fmt.Printf("%v", l.Val)
		l = l.Next
	}
}

func TestMergeTwoLists(t *testing.T) {
	for _, ut := range ucTests {
		fmt.Printf("testcase xxxxxxxxxxx\n")
		printList(ut.l1)
		uc := mergeTwoLists(ut.l1, ut.l2)

		if !isEqual(uc, ut.expected) {
			t.Errorf("l1 = %+v, l2 = %+v, got = %+v, must be %+v", ut.l1, ut.l2, uc, ut.expected)
			printList(uc)
			printList(ut.expected)
		}
	}
}

func isEqual(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 != nil {
		return false
	}
	if l1 != nil && l2 == nil {
		return false
	}
	if l1 == nil && l2 == nil {
		return true
	}

	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil && l2 == nil {
			return false
		}
		if l1 == nil && l2 != nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
