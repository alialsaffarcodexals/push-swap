package ps

// IsSorted checks whether the slice is sorted in ascending order.
func IsSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}
