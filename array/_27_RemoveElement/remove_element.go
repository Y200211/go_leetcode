package _27_RemoveElement

import "fmt"

func RemoveElement(nums []int, val int) (res []int, length int) {
	i, j := 0, 0
	for i = 0; i < len(nums); i++ {

		for j < len(nums) && nums[j] == val {
			j++
		}
		if j >= len(nums) {
			break
		}
		nums[i] = nums[j]
		j++
	}
	return nums, i
}

func TestMain() {
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res, length := RemoveElement(nums1, val)
	fmt.Println(res, length)
}
