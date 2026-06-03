func twoSum(nums []int, target int) []int {
    checked := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        num := nums[i]
        diff := target - num
        value, exists := checked[diff]
        if exists {
            return []int{value, i}
        }

        checked[num] = i
    }

    return []int{}
}
