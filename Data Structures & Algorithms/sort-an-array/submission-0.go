func sortArray(nums []int) []int {
   countSort(nums)
   return nums


}

func countSort(nums []int){
	 count:=make(map[int]int)
	 min,max:=nums[0],nums[0]
	 // Loop through entire nums and count each val, also maintain min and max
	 for _,v:=range nums{
		count[v]++
		if v<min{
			min=v
		}
		if v>max{
			max=v
		}
	 }
	 //Now update nums based on the count
	 i:=0
	 for v:=min;v<=max;v++{
		for count[v]>0{
			nums[i]=v
			i++
			count[v]--
		}
	 }
}
