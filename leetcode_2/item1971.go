package leetcode_2

type Union []int

func NewUnion(l int) Union {
	union := make([]int, l)
	for i := 0; i < l; i++ {
		union[i] = i
	}
	return union
}

func (r Union) Find(node int) int {
	for {
		parent := r[node]
		if parent == r[parent] {
			return parent
		}
		r[node] = r[parent]
	}
}
func (r Union) Connect(node1, node2 int) {
	p1 := r.Find(node1)
	p2 := r.Find(node2)
	if p1 == p2 {
		return
	}
	r[p1] = p2
}
func (r Union) IsConnect(node1, node2 int) bool {
	return r.Find(node1) == r.Find(node2)
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	union := NewUnion(n)
	for _, edge := range edges {
		union.Connect(edge[0], edge[1])
	}
	return union.IsConnect(source, destination)
}
func validPath1(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	adj := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		adj[edge[0]][edge[1]] = struct{}{}
		adj[edge[1]][edge[0]] = struct{}{}
	}
	q := make([]int, 0, n)
	q = append(q, source)
	seen := make([]bool, n)
	seen[source] = true
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for next := range adj[node] {
			if seen[next] {
				continue
			}
			seen[next] = true
			if next == destination {
				return true
			}
			q = append(q, next)
		}
	}
	return false
}
