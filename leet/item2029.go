package leet

func stoneGameIX(stones []int) bool {
	c := [3]int{}
	for _, v := range stones {
		c[v%3]++
	}
	if c[0]%2 == 0 {
		return c[1] > 0 && c[2] > 0
	}
	return c[1]-2 > c[2] || c[2]-2 > c[1]
}
