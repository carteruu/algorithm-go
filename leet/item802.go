package leet

import "sort"

func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	rg := make([][]int, n)
	inDeg := make([]int, n)
	for x, ys := range graph {
		for _, y := range ys {
			rg[y] = append(rg[y], x)
		}
		inDeg[x] = len(ys)
	}

	q := []int{}
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, x := range rg[y] {
			inDeg[x]--
			if inDeg[x] == 0 {
				q = append(q, x)
			}
		}
	}

	for i, d := range inDeg {
		if d == 0 {
			ans = append(ans, i)
		}
	}
	return
}

func eventualSafeNodes2(graph [][]int) []int {
	n := len(graph)
	color := make([]int, n)
	var isSalf func(node int) bool
	isSalf = func(node int) bool {
		if color[node] > 0 {
			return color[node] == 2
		}
		color[node] = 1
		for _, next := range graph[node] {
			if !isSalf(next) {
				return false
			}
		}
		color[node] = 2
		return true
	}
	var res []int
	for i := 0; i < n; i++ {
		if isSalf(i) {
			res = append(res, i)
		}
	}
	return res
}
func eventualSafeNodes1(graph [][]int) []int {
	n := len(graph)
	type idxLen struct {
		idx int
		l   int
	}
	ls := make([]*idxLen, n)
	for i := 0; i < n; i++ {
		ls[i] = &idxLen{i, len(graph[i])}
	}
	sort.Slice(ls, func(i, j int) bool {
		return ls[i].l < ls[j].l
	})

	choice := make([]bool, n)
	dead := make([]bool, n)
	target := make([]bool, n)
	var ans []int
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if target[idx] {
			return true
		}
		if dead[idx] {
			return false
		}
		for j := 0; j < len(graph[idx]); j++ {
			if choice[graph[idx][j]] {
				return false
			}
			choice[graph[idx][j]] = true
			if !dfs(graph[idx][j]) {
				return false
			}
			choice[graph[idx][j]] = false
		}
		return true
	}

	for i := 0; i < n; i++ {
		idx := ls[i].idx
		if dfs(idx) {
			ans = append(ans, idx)
			target[idx] = true
		} else {
			dead[idx] = true
		}
	}
	sort.Ints(ans)
	return ans
}
