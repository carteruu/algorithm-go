package inte

func majorityElement(nums []int) int {
	i := 0
	for i < len(nums) {
		j := i + 1
		for j < len(nums) {
			if nums[j] != nums[i] {
				break
			}
			j++
		}
		if j == len(nums) {
			return nums[i]
		}
		i++
	}
	return -1
}
