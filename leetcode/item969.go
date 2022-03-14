package leetcode

func pancakeSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	var ans []int
	end := len(arr)
	for end > 1 {
		maxIdx := 0
		max := arr[0]
		for i := 1; i < end; i++ {
			if arr[i] > max {
				maxIdx = i
				max = arr[i]
			}
		}
		if maxIdx != end {
			ans = append(ans, maxIdx+1, end)
			Reverse(arr[:maxIdx+1])
			Reverse(arr[:end])
		}
		end--
	}
	return ans
}
func Reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
