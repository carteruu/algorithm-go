package leet

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}
func findCenter1(edges [][]int) int {
	cnt := make([]int, len(edges)+2)
	for _, edge := range edges {
		cnt[edge[0]]++
		cnt[edge[1]]++
	}
	last := edges[len(edges)-1]
	if cnt[last[0]] == len(edges) {
		return last[0]
	} else {
		return last[1]
	}
}
