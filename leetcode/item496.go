package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//单调栈,存数值
	stack := make([]int, 0, len(nums2))
	//map:每个数字->下一个更大数字
	m := make(map[int]int, len(nums2))
	//遍历nums2,计算下一个更大数值
	for _, num := range nums2 {
		//更新答案并出栈的条件:当栈不为空,且栈顶的数值比当前数值小,得到栈顶的下一个更大数值,并出栈
		for l := len(stack); l > 0 && num > stack[l-1]; l = len(stack) {
			m[stack[l-1]] = num
			stack = stack[:l-1]
		}
		//将每个数字都加入栈
		stack = append(stack, num)
	}
	//得到答案
	ans := make([]int, len(nums1))
	for i, num1 := range nums1 {
		if val, ok := m[num1]; ok {
			ans[i] = val
		} else {
			ans[i] = -1
		}
	}
	return ans
}
