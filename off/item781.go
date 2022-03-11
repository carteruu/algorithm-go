package off

import "sort"

func numRabbits2(answers []int) int {
	count := map[int]int{}
	ans := 0
	for _, y := range answers {
		count[y]++
	}
	for y, x := range count {
		//ans += (x + y) / (y + 1) * (y + 1)
		c := x / (y + 1)
		if x%(y+1) > 0 {
			c++
		}
		ans += c * (y + 1)
	}
	return ans
}

func numRabbits(answers []int) int {
	if len(answers) == 0 {
		return 0
	}
	//排序
	sort.Ints(answers)
	res := answers[0] + 1
	//回答相同的兔子数
	same := 1
	for i := 1; i < len(answers); i++ {
		if answers[i] == answers[i-1] && same <= answers[i] {
			//回答和上一只兔子相同,
			//且相同回答的兔子数,小于等于回答的数量时,
			//可以认为相同回答的兔子颜色相同
			same++
			continue
		}
		//增加一批兔子
		res += answers[i] + 1
		same = 1
	}
	return res
}
