package binary_search

func Search(nums []int, target int) int {
	lowIdx, highIdx := 0, len(nums)-1
	for lowIdx <= highIdx {
		midIdx := (highIdx + lowIdx) / 2
		if nums[midIdx] == target {
			return midIdx
		}

		if nums[midIdx] < target {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx - 1
		}
	}

	return -1
}
