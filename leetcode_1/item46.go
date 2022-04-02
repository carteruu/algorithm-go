package leetcode_1

func permute(nums []int) [][]int {
	l := 1
	n := len(nums)
	for i := 1; i <= n; i++ {
		l *= i
	}
	ans := make([][]int, 0, l)
	/**
	map: 数字 -> 所在位置
	*/
	var dfs func(path map[int]int)
	dfs = func(path map[int]int) {
		if len(path) == n {
			item := make([]int, n)
			//将数字放到所在位置
			for key, val := range path {
				item[val] = key
			}
			ans = append(ans, item)
			return
		}
		for _, num := range nums {
			if _, ok := path[num]; ok {
				continue
			}
			path[num] = len(path)
			dfs(path)
			delete(path, num)
		}
	}
	dfs(make(map[int]int, n))
	return ans
}
func permute_1(nums []int) [][]int {
	l := 1
	n := len(nums)
	for i := 1; i <= n; i++ {
		l *= i
	}
	ans := make([][]int, 0, l)
	var dfs func(path []int, seen []bool)
	dfs = func(path []int, seen []bool) {
		if len(path) == n {
			item := make([]int, n)
			copy(item, path)
			ans = append(ans, item)
			return
		}
		for i, num := range nums {
			if seen[i] {
				continue
			}
			seen[i] = true
			path = append(path, num)
			dfs(path, seen)
			path = path[:len(path)-1]
			seen[i] = false
		}
	}
	dfs(make([]int, 0, n), make([]bool, n))
	return ans
}
