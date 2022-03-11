package item1506

type Node struct {
	Val      int
	Children []*Node
}

func findRoot(tree []*Node) *Node {
	a := 0
	for _, node := range tree {
		a ^= node.Val
		for _, child := range node.Children {
			a ^= child.Val
		}
	}
	for _, node := range tree {
		if node.Val == a {
			return node
		}
	}
	return nil
}
func findRoot1(tree []*Node) *Node {
	m := make(map[*Node]int)
	for _, node := range tree {
		if _, exist := m[node]; !exist {
			m[node] = 0
		}
		for _, child := range node.Children {
			m[child]++
		}
	}
	for node, times := range m {
		if times == 0 {
			return node
		}
	}
	return nil
}
