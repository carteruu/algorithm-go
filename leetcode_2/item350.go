package leetcode_2

func intersect(nums1 []int, nums2 []int) []int {
	cnt := make([]int, 1001)
	for _, num := range nums1 {
		cnt[num]++
	}
	var ans []int
	for _, num := range nums2 {
		if cnt[num] <= 0 {
			continue
		}
		cnt[num]--
		ans = append(ans, num)
	}
	return ans
}
