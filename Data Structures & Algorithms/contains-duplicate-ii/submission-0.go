func containsNearbyDuplicate(nums []int, k int) bool {
	hmap:=make(map[int]bool)
	l:=0

	for r:=0;r<len(nums);r++{
		if r-l>k{
			delete(hmap, nums[l])
			l++
		}
		if hmap[nums[r]]{
			return true
		}
		hmap[nums[r]]=true
	}
	return false
}
