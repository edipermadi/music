package util

// RotateLeftSliceInt rotate integer slice to the left k times
func RotateLeftSliceInt(nums []int, k int) []int {
	length := len(nums)
	if k < 0 || length == 0 {
		return nums
	}

	r := k % length
	return append(nums[k:], nums[:r]...)
}
