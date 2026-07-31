class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_so_far = {}

        for i in range(len(nums)):
            diff = target - nums[i]

            if diff in nums_so_far:
                return [nums_so_far[diff], i]
            
            nums_so_far[nums[i]] = i
        return []