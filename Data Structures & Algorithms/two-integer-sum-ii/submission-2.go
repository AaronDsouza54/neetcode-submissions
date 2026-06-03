func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	
	begin := 0
	end := len(numbers) - 1

	for numbers[begin] + numbers[end] != target {
		temp := numbers[begin] + numbers[end]
		if temp < target {
			begin++
		} else if temp > target {
			end--
		}
	}

	return []int{begin + 1, end + 1}
}
