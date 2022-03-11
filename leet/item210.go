package leet

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	inDeg := make([]int, numCourses)
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		inDeg[pre[0]]++
	}
	var q []int
	for i, deg := range inDeg {
		if deg == 0 {
			q = append(q, i)
		}
	}
	result := make([]int, 0, numCourses)
	for len(q) > 0 {
		cur := q[len(q)-1]
		q = q[:len(q)-1]
		result = append(result, cur)
		for _, next := range edges[cur] {
			inDeg[next]--
			if inDeg[next] == 0 {
				q = append(q, next)
			}
		}
	}
	if len(result) == numCourses {
		return result
	} else {
		return []int{}
	}
}
func findOrder1(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	result := make([]int, 0, numCourses)
	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node] = 1
		for _, next := range edges[node] {
			if visited[next] == 1 {
				return false
			} else if visited[next] == 0 {
				if !dfs(next) {
					return false
				}
			}
		}
		visited[node] = 2
		result = append(result, node)
		return true
	}

	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfs(i) {
				return []int{}
			}
		}
	}
	for i := 0; i < numCourses/2; i++ {
		result[i], result[numCourses-1-i] = result[numCourses-1-i], result[i]
	}
	return result
}
