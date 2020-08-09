/*
 * author: kaku
 * date: 2020/08/09
 *
 * GitHub:
 *   https://github.com/kakukosaku
 *
 * Description:
 *   leetcode #53
 */
class Solution {
    public int maxSubArray(int[] nums) {
        int maxSum = nums[0];
        int totalSum = 0;
        for (int val : nums) {
            totalSum = Math.max(val, totalSum + val);
            maxSum = Math.max(totalSum, maxSum);
        }

        return maxSum;
    }
}