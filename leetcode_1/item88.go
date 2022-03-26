package leetcode_1

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx, idx1, idx2 := m+n-1, m-1, n-1
	for idx1 >= 0 && idx2 >= 0 {
		if nums1[idx1] > nums2[idx2] {
			nums1[idx] = nums1[idx1]
			idx1--
		} else {
			nums1[idx] = nums2[idx2]
			idx2--
		}
		idx--
	}
	copy(nums1[:idx+1], nums1[:idx1+1])
	copy(nums1[:idx+1], nums2[:idx2+1])
}
