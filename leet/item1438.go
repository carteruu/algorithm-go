package leet

/*
创建一个队列
将新元素加入前,先和队尾元素比较:
	递减队列: 队尾元素较小,则删除队尾元素,保持单调性,队头为最大值;
	递增队列: 队尾元素较大,则删除队尾元素,保持单调性,队头为最小值;
将新元素加入队尾
*/
func longestSubarray(nums []int, limit int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	ans := 0
	//滑动窗口中的最大值和最小值
	qMin := make([]int, 0, len(nums))
	qMax := make([]int, 0, len(nums))
	for l, r := 0, 0; r < len(nums); r++ {
		//向右移动右边界,将元素加入窗口.维护最小值和最大值的单调队列
		for size := len(qMin); size > 0 && qMin[size-1] > nums[r]; size = len(qMin) {
			qMin = qMin[:size-1]
		}
		qMin = append(qMin, nums[r])
		for size := len(qMax); size > 0 && qMax[size-1] < nums[r]; size = len(qMax) {
			qMax = qMax[:size-1]
		}
		qMax = append(qMax, nums[r])
		//窗口不满足条件时,向右移动左边界,将元素移出窗口
		for l < r && (abs(nums[r]-qMin[0]) > limit || abs(qMax[0]-nums[r]) > limit) {
			//更新最小值和最大值的单调队列
			if nums[l] == qMin[0] {
				qMin = qMin[1:]
			}
			if nums[l] == qMax[0] {
				qMax = qMax[1:]
			}
			l++
		}
		//计算最大长度
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}
