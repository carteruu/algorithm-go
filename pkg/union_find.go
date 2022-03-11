package pkg

type UnionFind struct {
	parent []int
	size   []int
	cnt    int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
		cnt:    n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}
func (uf *UnionFind) Find(node int) int {
	parent := uf.parent[node]
	for uf.parent[parent] != parent {
		uf.parent[node] = uf.parent[parent]
		uf.size[parent] -= uf.size[node]
		parent = uf.parent[node]
	}
	return parent
}
func (uf *UnionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)
	if rootA == rootB {
		return false
	}
	var weight, light int
	if uf.size[rootA] < uf.size[rootB] {
		weight, light = rootB, rootA
	} else {
		weight, light = rootA, rootB
	}
	uf.parent[light] = weight
	uf.size[weight] += uf.size[light]
	//连通分量-1
	uf.cnt--
	return true
}
func (uf *UnionFind) Connected(a, b int) bool {
	return uf.Find(a) == uf.Find(b)
}
func (uf *UnionFind) Size(node int) int {
	return uf.size[node]
}
