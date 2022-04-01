package leetcode_1

/**
并查集
*/
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	//只需要和右、下两个方向连接
	steps := [][]int{{0, 1}, {1, 0}}
	uf := NewUnionFind(m*n + 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			node := i*n + j
			if grid[i][j] == '0' {
				//水域，都和一个特殊的节点连接，最后连通分量减一即是答案
				uf.Union(node, m*n)
				continue
			}
			for _, step := range steps {
				x, y := i+step[0], j+step[1]
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == '1' {
					uf.Union(node, x*n+y)
				}
			}
		}
	}
	return uf.Cnt() - 1
}
func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	return &UnionFind{
		parents: parents,
		cnt:     n,
	}
}

type UnionFind struct {
	parents []int
	//连通分量
	cnt int
}

func (uf *UnionFind) Union(node1, node2 int) bool {
	if uf.IsConnect(node1, node2) {
		return false
	}
	parent1 := uf.Find(node1)
	parent2 := uf.Find(node2)
	uf.parents[parent1] = parent2
	uf.cnt--
	return true
}
func (uf *UnionFind) Find(node int) int {
	//如果父节点的父不是父节点自身（父节点非根），则将当前节点移动到祖父节点，以降低树的高度
	for uf.parents[uf.parents[node]] != uf.parents[node] {
		uf.parents[node] = uf.parents[uf.parents[node]]
	}
	return uf.parents[node]
}
func (uf *UnionFind) IsConnect(node1, node2 int) bool {
	return uf.Find(node1) == uf.Find(node2)
}
func (uf UnionFind) Cnt() int {
	return uf.cnt
}
