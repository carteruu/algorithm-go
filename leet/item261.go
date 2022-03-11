package leet

func validTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}
	if len(edges) != n-1 {
		return false
	}
	//邻接表-双向
	adjList := make([]map[int]struct{}, n)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		adjList[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		adjList[edge[0]][edge[1]] = struct{}{}
		adjList[edge[1]][edge[0]] = struct{}{}
	}
	//深度优先
	var dfs func(node int) bool
	dfs = func(node int) bool {
		//标记当前节点已访问
		visited[node] = true
		//遍历子节点
		for k := range adjList[node] {
			if visited[k] {
				//子节点已经被访问过了,存在环
				return false
			} else {
				//删除子节点到当前节点的边
				delete(adjList[k], node)
				if !dfs(k) {
					return false
				}
			}
		}
		return true
	}
	if !dfs(0) {
		return false
	}
	//判断是否所有节点都被访问过
	for i := 0; i < n; i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}
func validTree2(n int, edges [][]int) bool {
	adjList := make([]map[int]struct{}, n)
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		adjList[edge[0]][edge[1]] = struct{}{}
		adjList[edge[1]][edge[0]] = struct{}{}
	}

	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node] = 1
		for k := range adjList[node] {
			if visited[k] == 1 {
				return false
			} else if visited[k] == 0 {
				delete(adjList[k], node)
				if !dfs(k) {
					return false
				}
			}
		}
		visited[node] = 2
		return true
	}
	if !dfs(0) {
		return false
	}
	for i := 0; i < n; i++ {
		if visited[i] != 2 {
			return false
		}
	}
	return true
}
func validTree1(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}
	//邻接表-双向
	adjList := make([]map[int]struct{}, n)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		adjList[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		adjList[edge[0]][edge[1]] = struct{}{}
		adjList[edge[1]][edge[0]] = struct{}{}
	}
	//将任意一个节点作为根节点,需要能访问到所有的节点,且不能存在环.
	//这里选择第一个节点,使用广度优先遍历
	queue := make([]int, 0, n)
	queue = append(queue, 0)
	for len(queue) > 0 {
		cur := queue[0]
		if visited[cur] {
			//如果当前节点已经被访问过了,说明存在环,不能构成树
			return false
		}
		queue = queue[1:]
		for next := range adjList[cur] {
			//依次遍历子节点,删除子节点到当前节点的边,并加入队列
			//没必要删除当前节点到子节点的边,因为之后不再需要用当前节点遍历了
			delete(adjList[next], cur)
			queue = append(queue, next)
		}
		//标记当前节点已访问
		visited[cur] = true
	}
	//判断是否全部节点都被访问
	for i := 0; i < n; i++ {
		if !visited[i] {
			//否,不满足条件
			return false
		}
	}
	return true
}

func validTree3(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	count := n
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for i := 0; i < len(edges); i++ {
		if union(edges[i][0], edges[i][1], parent) {
			return false
		}
		count--
	}
	return count == 1
}

func union(x, y int, parent []int) bool {
	fx := find(x, parent)
	fy := find(y, parent)
	if fx == fy {
		return true
	}
	parent[fx] = fy
	return false
}

func find(x int, parent []int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x], parent)
	return parent[x]
}
