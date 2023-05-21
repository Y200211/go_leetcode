package _704_BinarySearch

func Binary_search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	n := len(nums) - 1
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
