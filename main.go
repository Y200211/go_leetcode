package main

import (
	"fmt"
	"go_leetcode/array/_27_RemoveElement"
)

func main() {
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res, length := _27_RemoveElement.RemoveElement(nums1, val)
	fmt.Println(res, length)
}
