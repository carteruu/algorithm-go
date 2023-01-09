package leetcode_2

import "sort"

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool { return edgeList[i][2] < edgeList[j][2] })
	// 并查集模板
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		fa[find(from)] = find(to)
	}

	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	// 按照 limit 从小到大排序，方便离线
	sort.Slice(queries, func(i, j int) bool { return queries[i][2] < queries[j][2] })

	ans := make([]bool, len(queries))
	k := 0
	for _, q := range queries {
		for ; k < len(edgeList) && edgeList[k][2] < q[2]; k++ {
			merge(edgeList[k][0], edgeList[k][1])
		}
		ans[q[3]] = find(q[0]) == find(q[1])
	}
	return ans
}
func distanceLimitedPathsExist2(n int, edgeList [][]int, queries [][]int) []bool {
	adjs := make(map[int]map[int]int, n)
	for _, edge := range edgeList {
		if adjs[edge[0]] == nil {
			adjs[edge[0]] = make(map[int]int)
		}
		if adjs[edge[1]] == nil {
			adjs[edge[1]] = make(map[int]int)
		}
		if d, ok := adjs[edge[0]][edge[1]]; !ok || edge[2] < d {
			adjs[edge[0]][edge[1]] = edge[2]
			adjs[edge[1]][edge[0]] = edge[2]
		}
	}
	ans := make([]bool, len(queries))
	for i, query := range queries {
		seen := make(map[int]struct{})
		var q []int
		q = append(q, query[0])
		for len(q) > 0 {
			node := q[0]
			if node == query[1] {
				ans[i] = true
				break
			}
			q = q[1:]
			seen[node] = struct{}{}
			for next, d := range adjs[node] {
				if _, ok := seen[next]; ok {
					continue
				}
				if d >= query[2] {
					continue
				}
				q = append(q, next)
			}
		}
	}
	return ans
}
func distanceLimitedPathsExist1(n int, edgeList [][]int, queries [][]int) []bool {
	adjs := make([][][2]int, n)
	for _, edge := range edgeList {
		adjs[edge[0]] = append(adjs[edge[0]], [2]int{edge[1], edge[2]})
		adjs[edge[1]] = append(adjs[edge[1]], [2]int{edge[0], edge[2]})
	}
	ans := make([]bool, len(queries))
	for i, query := range queries {
		seen := make(map[int]struct{})
		var q []int
		q = append(q, query[0])
		for len(q) > 0 {
			node := q[0]
			if node == query[1] {
				ans[i] = true
				break
			}
			q = q[1:]
			seen[node] = struct{}{}
			for _, next := range adjs[node] {
				if _, ok := seen[next[0]]; ok {
					continue
				}
				if next[1] >= query[2] {
					continue
				}
				q = append(q, next[0])
			}
		}
	}
	return ans
}
