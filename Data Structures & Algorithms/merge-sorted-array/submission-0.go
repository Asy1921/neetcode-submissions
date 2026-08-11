func merge(nums1 []int, m int, nums2 []int, n int) {
	lastIdx := m + n - 1
	i, j := m-1, n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[lastIdx] = nums1[i]
			i--
		} else {
			nums1[lastIdx] = nums2[j]
			j--
		}
		lastIdx--
	}
}