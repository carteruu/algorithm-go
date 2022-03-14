package leetcode

func candy(ratings []int) int {
	if ratings == nil || len(ratings) == 0 {
		return 0
	}
	n := len(ratings)
	cs := make([]int, n)
	//从左往右
	//第一个小孩
	cs[0] = 1
	//从第二个小孩开始遍历
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			cs[i] = cs[i-1] + 1
		} else {
			cs[i] = 1
		}
	}

	//从右往左
	//第一个小孩
	val := max(1, cs[n-1])
	ans := val
	//从第二个小孩开始遍历
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			val += 1
		} else {
			val = 1
		}
		//取从左往右和从右往左计算出的值的较大值
		ans += max(val, cs[i])
	}
	return ans
}
func candy1(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

func candy3(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
