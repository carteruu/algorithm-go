package leetcode

import (
	"math"
	"strconv"
)

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	if n <= 1 {
		return 0
	}
	seen := make([][]bool, n)
	type node struct {
		idx     int
		visited int
		dist    int
	}

	q := make([]*node, n)
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, 1<<n)
		seen[i][i] = true
		q[i] = &node{i, 1 << i, 0}
	}
	t := 1<<n - 1
	for {
		poll := q[0]
		q = q[1:]
		if poll.visited == t {
			return poll.dist
		}
		for _, next := range graph[poll.idx] {
			visitedN := poll.visited | 1<<next
			if !seen[next][visitedN] {
				seen[next][visitedN] = true
				q = append(q, &node{next, visitedN, poll.dist + 1})
			}
		}
	}
}

func shortestPathLength2(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, tuple{i, 1 << i, 0})
	}

	for {
		t := q[0]
		q = q[1:]
		if t.mask == 1<<n-1 {
			return t.dist
		}
		// 搜索相邻的节点
		for _, v := range graph[t.u] {
			maskV := t.mask | 1<<v
			if !seen[v][maskV] {
				q = append(q, tuple{v, maskV, t.dist + 1})
				seen[v][maskV] = true
			}
		}
	}
}
func shortestPathLength1(graph [][]int) int {
	visited := make([]bool, len(graph))
	visitedCount := 1
	path := make(map[string]struct{})
	res := math.MaxInt64

	var dfs func(node, router int)
	dfs = func(node, router int) {
		if visitedCount == len(graph) {
			if router < res {
				res = router
			}
		}
		for i := 0; i < len(graph[node]); i++ {
			next := graph[node][i]
			key := strconv.Itoa(node) + "-" + strconv.Itoa(next)
			if _, ok := path[key]; ok {
				continue
			}
			path[key] = struct{}{}
			if visited[next] {
				//之前已访问过
				dfs(next, router+1)
			} else {
				visited[next] = true
				visitedCount++
				dfs(next, router+1)
				visited[next] = false
				visitedCount--
			}
			delete(path, key)
		}
	}

	for i := 0; i < len(graph); i++ {
		visited[i] = true
		dfs(i, 0)
		visited[i] = false
	}
	return res
}
