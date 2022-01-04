package remove_duplicates_from_sorted_array

//removeDuplicates Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that
// each unique element appears only once.
func removeDuplicates(nums []int) int {

	// 如果数组的长度为 0，则数组不包含任何元素，因此返回 0
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 双指针
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
