func getConcatenation(nums []int) []int {
    ans := make([]int, 2*len(nums))

	n := len(nums)

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		ans[i+n] = nums[i]
	}

	return ans
}
