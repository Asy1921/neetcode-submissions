public class Solution {
    public int[] GetConcatenation(int[] nums) {
        int length=nums.Length;
        int n=2*length;
        int[] resultArray=new int[n];
        for(int i=0;i<n;i++)
        {   int p=i;
            if(i>=length)
            {
                p=i-length;
            }
            resultArray[i]=nums[p];
        }
        return resultArray;
    }
}