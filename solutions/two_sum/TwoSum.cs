public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        for(int fixedNum = 0; fixedNum < nums.Length; fixedNum++)
        {
            for (int i = 0; i < nums.Length; i++)
            {
                if ((fixedNum != i) && (nums[fixedNum] + nums[i] == target))
                {
                    return new int[] { fixedNum, i };
                }
            }
        }
        return null;
    }
}