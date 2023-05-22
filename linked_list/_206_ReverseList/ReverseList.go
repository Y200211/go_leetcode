package _206_ReverseList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	cur := head.Next
	next := head.Next.Next
	prev.Next = nil
	if next == nil {
		cur.Next = prev
	}

	for next != nil {
		tmp := next.Next
		cur.Next = prev
		next.Next = cur
		prev = cur
		cur = next
		next = tmp
	}
	return cur
}

func CreateLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{Val: nums[i]}
		tail = tail.Next
	}
	return head
}
func ShowLinkedList(head *ListNode) {
	if head == nil {
		return
	}
	tail := head
	for tail != nil {
		fmt.Println(tail.Val)
		tail = tail.Next
	}
}
