package leetcode_2

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
