func getConcatenation(nums []int) []int {
    	combined := make([]int, len(nums)*2)
	copy(combined, nums)
	copy(combined[len(nums):], nums)
	return combined
}
