package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Next == nil && l2.Next == nil {
		carry := (l1.Val + l2.Val) / 10
		val := (l1.Val + l2.Val) % 10
		if carry != 0 {
			return &ListNode{Val: val, Next: &ListNode{Val: carry, Next: nil}}
		}
		return &ListNode{Val: val, Next: nil}
	} else if l1.Next == nil {
		carry := (l1.Val + l2.Val) / 10
		val := (l1.Val + l2.Val) % 10
		l2.Next.Val = l2.Next.Val + carry
		return &ListNode{Val: val, Next: addTwoNumbers(&ListNode{Val: 0, Next: nil}, l2.Next)}
	} else if l2.Next == nil {
		carry := (l1.Val + l2.Val) / 10
		val := (l1.Val + l2.Val) % 10
		l1.Next.Val = l1.Next.Val + carry
		return &ListNode{Val: val, Next: addTwoNumbers(l1.Next, &ListNode{Val: 0, Next: nil})}
	} else {
		carry := (l1.Val + l2.Val) / 10
		val := (l1.Val + l2.Val) % 10
		l1.Next.Val = l1.Next.Val + carry
		return &ListNode{Val: val, Next: addTwoNumbers(l1.Next, l2.Next)}
	}
}

// l1 = [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
// l2 = [5,6,4]
// result = [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]

func main() {
	final := addTwoNumbers(testasseValue(), testasseValue2())
	printList(final)
}

func testasseValue() *ListNode {
	return &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
}

func testasseValue2() *ListNode {
	return &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
}

func printList(l *ListNode) {
	for l.Next != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
}
