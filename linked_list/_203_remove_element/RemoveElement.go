package _203_remove_element

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func RemoveElement(head *Node, num int) *Node {

	if head == nil {
		return nil
	}
	for head != nil && head.Val == num {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	tail := head
	for tail.Next != nil {
		for tail.Next != nil && tail.Next.Val == num {
			tail.Next = tail.Next.Next
		}
		if tail.Next != nil {
			tail = tail.Next
		} else {
			break
		}

	}
	return head
}

func CreateLinkedList(nums []int) *Node {
	head := &Node{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &Node{
			Val: nums[i],
		}
		tail = tail.Next
	}
	return head
}

func PrintLinkedList(head *Node, length int) {
	fmt.Println(head.Val)
	for head.Next != nil {
		head = head.Next
		fmt.Println(head.Val)
	}
}
