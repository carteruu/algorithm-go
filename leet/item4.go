package leet

func findMedianSortedArrays11(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m+n == 0 {
		return 0
	}
	nums := make([]int, m+n)
	s1, s2 := 0, 0
	for i := 0; i < m+n; i++ {
		if s2 == n || (s1 < m && nums1[s1] < nums2[s2]) {
			nums[i] = nums1[s1]
			s1++
		} else {
			nums[i] = nums2[s2]
			s2++
		}
	}
	if (m+n)&1 == 1 {
		//奇数
		return float64(nums[(m+n)/2])
	} else {
		return float64(nums[(m+n)/2-1]+nums[(m+n)/2]) / 2
	}
}
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	l := m + n
	if l == 0 {
		return 0
	}
	s1, s2 := 0, 0
	var pre, cur int
	for i := 0; i <= l/2; i++ {
		pre = cur
		if s2 == n || (s1 < m && nums1[s1] < nums2[s2]) {
			cur = nums1[s1]
			s1++
		} else {
			cur = nums2[s2]
			s2++
		}
	}
	if l&1 == 1 {
		//奇数
		return float64(cur)
	} else {
		//偶数
		return float64(pre+cur) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}
