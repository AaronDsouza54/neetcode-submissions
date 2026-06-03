func hasDuplicate(nums []int) bool {
    temp := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		_, exists := temp[num]
		
		if exists {
			return true
		}

		temp[num] = struct{}{}
	}

	return false
}
