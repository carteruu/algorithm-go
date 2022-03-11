package leet

func dailyTemperatures1(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperature {
				ans[i] = j - i
				break
			}
		}
	}
	return ans
}

//从后往前遍历
func dailyTemperatures12(temperatures []int) []int {
	//单调栈
	stack := make([]int, 0, len(temperatures))
	ans := make([]int, len(temperatures))
	//从尾部开始遍历
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			//如果栈不为空,且栈顶的温度不高于今天温度,出栈.
			//因为题目求的是几天后的温度更高,后面不高于今天温度的气温已经不需要了.
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			//栈不为空,说明存在温度比今天高的
			ans[i] = stack[len(stack)-1] - i
		}
		//将今天温度加入栈顶.形成单调栈,从栈底到栈顶,温度递减
		stack = append(stack, i)
	}
	return ans
}

//从前往后遍历
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			//栈不为空,且今天温度大于栈顶日期的温度.更新栈顶日期的答案,并出栈
			prevIndex := stack[len(stack)-1]
			ans[prevIndex] = i - prevIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
