package main

// import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, carry := new(ListNode), 0
	for current := head; l1 != nil || l2 != nil || carry > 0; current = current.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{carry % 10, nil}
		carry /= 10
	}
	return head.Next
}

func printList(l *ListNode) {
	for l != nil {
		print(l.Val)
		l = l.Next
	}
	print("\n")
}

func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{Val: 3}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{Val: 4}}}
	printList(&l1)
	printList(&l2)
	printList(addTwoNumbers(&l1, &l2))

	l1 = ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	printList(&l1)
	printList(&l2)
	printList(addTwoNumbers(&l1, &l2))
}
