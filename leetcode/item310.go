package leetcode

import "math"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adjList := make([][]int, n)
	degs := make([]int, n)
	var queue, ans []int
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
		degs[edge[0]]++
		degs[edge[1]]++
	}
	for i := 0; i < n; i++ {
		if degs[i] == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		ans = make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			ans = append(ans, cur)
			for _, next := range adjList[cur] {
				degs[next]--
				if degs[next] == 1 {
					queue = append(queue, next)
				}
			}
		}
	}
	return ans
}

//超时了
func findMinHeightTrees1(n int, edges [][]int) []int {
	//邻接表-双向
	adjList := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		adjList[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		adjList[edge[0]][edge[1]] = struct{}{}
		adjList[edge[1]][edge[0]] = struct{}{}
	}

	not := make([]bool, n)

	var ans []int
	min := math.MaxInt32
	for i := 0; i < n; i++ {
		if not[i] {
			continue
		}
		visited := make([]bool, n)
		queue := make([]int, 1, n)
		queue[0] = i
		depth := 0
		for len(queue) > 0 {
			size := len(queue)
			for j := 0; j < size; j++ {
				cur := queue[0]
				queue = queue[1:]
				for next := range adjList[cur] {
					if visited[next] {
						continue
					}
					queue = append(queue, next)
				}
				visited[cur] = true
				if depth > min {
					not[cur] = true
				}
			}
			depth++
		}
		if depth < min {
			min = depth
			ans = []int{i}
		} else if depth == min {
			ans = append(ans, i)
		}
	}
	return ans
}

func findMinHeightTrees2(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	degrees := make([]int, n)
	for _, edge := range edges {
		degrees[edge[0]] += 1
		degrees[edge[1]] += 1
		if graph[edge[0]] == nil {
			graph[edge[0]] = []int{}
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		if graph[edge[1]] == nil {
			graph[edge[1]] = []int{}
		}
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	queue := []int{}
	for i := 0; i < n; i++ {
		if degrees[i] == 1 {
			queue = append(queue, i)
		}
	}
	roots := []int{}
	for len(queue) > 0 {
		length := len(queue)
		roots = []int{}
		for i := 0; i < length; i++ {
			degrees[queue[i]] -= 1
			for _, neighbor := range graph[queue[i]] {
				degrees[neighbor] -= 1
				if degrees[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
			roots = append(roots, queue[i])
		}
		queue = queue[length:]
	}
	return roots

}
