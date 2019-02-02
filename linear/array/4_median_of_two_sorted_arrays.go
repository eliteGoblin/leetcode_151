package array

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		return findKth(nums1, nums2, (totalLen-1)/2)
	}
	return (findKth(nums1, nums2, (totalLen)/2-1) +
		findKth(nums1, nums2, (totalLen)/2)) / 2.0
}

func findKth(nums1 []int, nums2 []int, k int) float64 {
	fmt.Printf("findKth %+v %+v k:%d\n", nums1, nums2, k)
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findKth(nums2, nums1, k)
	}
	if 0 == m {
		return float64(nums2[k])
	}
	if k == 0 {
		return float64(min(nums1[0], nums2[0]))
	}
	nums1PartLen := min((k+1)/2, m)
	nums2PartLen := k + 1 - nums1PartLen
	if nums1[nums1PartLen] == nums2[nums2PartLen] {
		return float64(nums1[nums1PartLen])
	} else if nums1[nums1PartLen] < nums2[nums1PartLen] {
		// 在nums1PartLen右侧
		return findKth(nums1[nums1PartLen:], nums2, k-nums1PartLen)
	}
	// nums1PartLen左侧
	return findKth(nums1, nums2[nums2PartLen:], k-nums2PartLen-1)
}
