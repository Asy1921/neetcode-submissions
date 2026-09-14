func subsets(nums []int) [][]int {
	res:= [][]int {}
	curr:=[]int{}

	var backtrack func(int)
	backtrack=func(i int){
		if i>=len(nums){
			//Add to result
			temp:=make([]int, len(curr))
			copy(temp,curr)
			res=append(res,temp)
			return
		}
		// Include current element i to curr
		curr=append(curr,nums[i])
		backtrack(i+1)
		// Dont include current element
		curr=curr[0:len(curr)-1]
		backtrack(i+1)
	}
	backtrack(0)
	return res
}
