package leetcode_2

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	return 1
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := 0
	l2 := 0
	numL, numR := 0, 0
	for k := 0; k <= (len(nums1)+len(nums2))/2; k++ {
		numL = numR
		if l1 < len(nums1) && (l2 >= len(nums2) || nums1[l1] < nums2[l2]) {
			numR = nums1[l1]
			l1++
		} else {
			numR = nums2[l2]
			l2++
		}
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(numL+numR) / 2
	} else {
		return float64(numR)
	}
}
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))
	var i, j int
	for k := 0; k < len(nums); k++ {
		if j == len(nums2) || (i < len(nums1) && nums1[i] < nums2[j]) {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}
	if l := len(nums); l%2 == 0 {
		return float64(nums[l/2-1]+nums[l/2]) / 2
	} else {
		return float64(nums[l/2])
	}
}
