func sortColors(nums []int) {
    count := make([]int,3)
	for _,v := range nums{
		count[v]++
	}
	idx:=0
	for ;idx<=2;idx++{
		if count[idx]>0{
			break;
		}
	}
	for i,_:=range nums{
		nums[i]=idx
		count[idx]--
		for ;idx<=2;idx++{
		if count[idx]>0{
			break;
		}
	}
	}
}
