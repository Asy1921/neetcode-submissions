func removeElement(nums []int, val int) int {
    	occurrences := 0
	pointer := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			occurrences++
		} else {
			nums[pointer] = nums[i]
			pointer++
		}
	}
	return len(nums) - occurrences
}
