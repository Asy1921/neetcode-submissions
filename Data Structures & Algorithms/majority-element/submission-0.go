func majorityElement(nums []int) int {
  res:=0
	count:=0

	for i:=0;i<len(nums);i++{
		if count==0{
			res=nums[i]
		}
		if nums[i]==res {
			count++
		} else {
			count--
		}
	}
	return res
}
