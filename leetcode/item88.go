package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1 := m - 1
	idx2 := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if idx2 < 0 || (idx1 >= 0 && nums1[idx1] > nums2[idx2]) {
			nums1[i] = nums1[idx1]
			idx1--
		} else {
			nums1[i] = nums2[idx2]
			idx2--
		}
	}
}
