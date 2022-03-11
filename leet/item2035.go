package leet

import (
	"math"
	"sort"
)

func minimumDifference(nums []int) int {
	n := len(nums) / 2
	a := nums[:n]
	res := make([][]int, n+1)
	for i := 0; i < 1<<n; i++ {
		sum, cnt := 0, 0
		for j, v := range a {
			if i>>j&1 > 0 { // 1 视作取正
				sum += v
				cnt++
			} else { // 0 视作取负
				sum -= v
			}
		}
		res[cnt] = append(res[cnt], sum) // 按照取正的个数将元素和分组
	}

	for _, b := range res {
		sort.Ints(b) // 排序，方便下面二分
	}

	ans := math.MaxInt64
	a = nums[n:]
	for i := 0; i < 1<<n; i++ {
		sum, cnt := 0, 0
		for j, v := range a {
			if i>>j&1 > 0 {
				sum += v
				cnt++
			} else {
				sum -= v
			}
		}
		// 在对应的组里二分最近的数
		b := res[cnt]
		j := sort.SearchInts(b, sum)
		if j < len(b) {
			ans = min(ans, b[j]-sum)
		}
		if j > 0 {
			ans = min(ans, sum-b[j-1])
		}
	}
	return ans
}
