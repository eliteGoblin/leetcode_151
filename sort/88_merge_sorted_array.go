package sort

func merge(nums1 []int, m int, nums2 []int, n int) {
	sortedIndex := m + n - 1
	nums1Index := m - 1
	nums2Index := n - 1
	for nums1Index >= 0 && nums2Index >= 0 {
		if nums1[nums1Index] >= nums2[nums2Index] {
			nums1[sortedIndex] = nums1[nums1Index]
			nums1Index--
		} else {
			nums1[sortedIndex] = nums2[nums2Index]
			nums2Index--
		}
		sortedIndex--
	}
	for nums1Index >= 0 {
		nums1[sortedIndex] = nums1[nums1Index]
		sortedIndex--
		nums1Index--
	}
	for nums2Index >= 0 {
		nums1[sortedIndex] = nums2[nums2Index]
		sortedIndex--
		nums2Index--
	}
}
