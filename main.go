package main

import (
	"fmt"
	"go_leetcode/linked_list/_206_ReverseList"
)

func main() {
	nums1 := []int{1, 2, 6, 3, 4, 5, 6}
	head1 := _206_ReverseList.CreateLinkedList(nums1)
	_206_ReverseList.ShowLinkedList(head1)
	fmt.Println("----------------")
	head2 := _206_ReverseList.ReverseList(head1)
	_206_ReverseList.ShowLinkedList(head2)
}
