package main

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
	var currentNode = &ListNode{}
	var carry = 0
	var result = currentNode
	for l1 != nil || l2 != nil {
		var val1 = 0
		var val2 = 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		var sum = val1 + val2 + carry
		carry = sum / 10
		currentNode.Val = sum % 10
		if l1 != nil || l2 != nil {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next
		} else {
			if carry > 0 {
				currentNode.Next = &ListNode{}
				currentNode = currentNode.Next
				currentNode.Val = carry
			}
		}
	}
	return result
}

func main() {
	list1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list := addTwoNumbers(list1, list2)
	for list != nil {
		println(list.Val)
		list = list.Next
	}
}
