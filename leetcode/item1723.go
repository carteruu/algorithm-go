package leetcode

import (
	"math"
	"sort"
)

func minimumTimeRequired2(jobs []int, k int) int {
	//sort.Ints(jobs)
	max := 0
	sum := 0
	for _, job := range jobs {
		if job > max {
			max = job
		}
		sum += job
	}

	var check func(t int) bool
	check = func(t int) bool {
		arr := make([]int, k)
		var helper func(jobIdx int) bool
		helper = func(jobIdx int) bool {
			if jobIdx == len(jobs) {
				return true
			}
			for i := 0; i < k; i++ {
				if arr[i]+jobs[jobIdx] <= t {
					arr[i] += jobs[jobIdx]
					if helper(jobIdx + 1) {
						return true
					}
					arr[i] -= jobs[jobIdx]
				}
				if arr[i] == 0 || arr[i]+jobs[jobIdx] == t {
					break
				}
			}
			return false
		}
		return helper(0)
	}

	for max < sum {
		mid := (sum-max)/2 + max
		if check(mid) {
			sum = mid
		} else {
			max = mid + 1
		}
	}
	return sum
}

func minimumTimeRequired1(jobs []int, k int) int {
	n := len(jobs)
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	l, r := jobs[0], 0
	for _, v := range jobs {
		r += v
	}
	return l + sort.Search(r-l, func(limit int) bool {
		limit += l
		workloads := make([]int, k)
		var backtrack func(int) bool
		backtrack = func(idx int) bool {
			if idx == n {
				return true
			}
			cur := jobs[idx]
			for i := range workloads {
				if workloads[i]+cur <= limit {
					workloads[i] += cur
					if backtrack(idx + 1) {
						return true
					}
					workloads[i] -= cur
				}
				// 如果当前工人未被分配工作，那么下一个工人也必然未被分配工作
				// 或者当前工作恰能使该工人的工作量达到了上限
				// 这两种情况下我们无需尝试继续分配工作
				if workloads[i] == 0 || workloads[i]+cur == limit {
					break
				}
			}
			return false
		}
		return backtrack(0)
	})
}

func minimumTimeRequired(jobs []int, k int) int {
	res := math.MaxInt32
	var dfs func(arr []int, jobIdx, maxTime, workerIdx int)
	dfs = func(arr []int, jobIdx, maxTime, workerIdx int) {
		if maxTime > res {
			return
		}
		if jobIdx == len(jobs) {
			res = maxTime
			return
		}
		//增加一个工人
		if workerIdx < k {
			arr[workerIdx] = jobs[jobIdx]
			dfs(arr, jobIdx+1, max(maxTime, arr[workerIdx]), workerIdx+1)
			arr[workerIdx] = 0
		}
		//分配给已有工作的工人
		for i := 0; i < workerIdx; i++ {
			arr[i] = arr[i] + jobs[jobIdx]
			dfs(arr, jobIdx+1, max(maxTime, arr[i]), workerIdx)
			arr[i] = arr[i] - jobs[jobIdx]
		}
	}
	dfs(make([]int, k), 0, 0, 0)
	return res
}
