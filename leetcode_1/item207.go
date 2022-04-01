package leetcode_1

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjs := make([][]int, numCourses)
	ins := make([]int, numCourses)
	for _, p := range prerequisites {
		adjs[p[1]] = append(adjs[p[1]], p[0])
		ins[p[0]]++
	}
	var q []int
	for i, cnt := range ins {
		if cnt == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		poll := q[len(q)-1]
		q = q[:len(q)-1]
		numCourses--
		for _, next := range adjs[poll] {
			ins[next]--
			if ins[next] == 0 {
				q = append(q, next)
			}
		}
	}
	return numCourses == 0
}
