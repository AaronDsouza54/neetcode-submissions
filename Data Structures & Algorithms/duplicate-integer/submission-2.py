class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        nums_so_far = {}

        for elem in nums:
            if elem in nums_so_far:
                return True
            nums_so_far[elem] = None
        return False