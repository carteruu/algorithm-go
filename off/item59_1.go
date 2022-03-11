package off

func maxSlidingWindow(nums []int, k int) []int {
	//双端队列
	q := make([]int, 0, len(nums))
	ans := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		//向右移动窗口
		for size := len(q); size > 0 && nums[i] > q[size-1]; size = len(q) {
			//当加入的元素比队尾的元素大时,删除队尾元素
			q = q[:size-1]
		}
		//加入队尾
		q = append(q, nums[i])

		if i >= k-1 {
			//窗口中的元素等于k时,计算答案
			ans = append(ans, q[0])
			//窗口左端的元素移出窗口
			if nums[i-k+1] == q[0] {
				//移出的元素等于队头时,队头出队
				q = q[1:]
			}
		}
	}
	return ans
}
