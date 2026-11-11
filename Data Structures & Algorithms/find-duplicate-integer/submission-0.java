class Solution {
    public int findDuplicate(int[] nums) {
        HashSet<Integer> seen = new HashSet<>();

        for (Integer val : nums) {
            if (seen.contains(val)) {
                return val;
            }

            seen.add(val);
        }

        return 0;
    }
}
