package remove

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	realLen := l
	for i := 1; i < realLen; {
		if nums[i] != nums[i-1] {
			i++
			continue
		}
		for j := i + 1; j < realLen; j++ {
			nums[j-1] = nums[j]
		}
		realLen = realLen - 1
	}

	return realLen
}
