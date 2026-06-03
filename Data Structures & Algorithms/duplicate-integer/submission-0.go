func hasDuplicate(nums []int) bool {
    temp := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		_, exists := temp[num]
		
		if exists {
			return true
		}

		temp[num] = true
	}

	return false
}
