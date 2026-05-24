package solution

import "testing"

func TestFindPeakElement(t *testing.T) {
	t.Skip("TODO: implement findPeakElement")
	tests := [][]int{
		[]int{1, 2, 3, 1},
		[]int{1, 2, 1, 3, 5, 6, 4},
		[]int{1},
	}
	for _, nums := range tests {
		i := findPeakElement(nums)
		if i < 0 || i >= len(nums) {
			t.Errorf("%v: index %d out of range", nums, i)
			continue
		}
		if i > 0 && nums[i] <= nums[i-1] {
			t.Errorf("%v: nums[%d]=%d not > left", nums, i, nums[i])
		}
		if i < len(nums)-1 && nums[i] <= nums[i+1] {
			t.Errorf("%v: nums[%d]=%d not > right", nums, i, nums[i])
		}
	}
}
