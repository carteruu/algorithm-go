package leetcode

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
func removeDuplicates1(nums []int) int {
	pre := nums[0]
	cnt := 1
	n := len(nums)
	sign := 10001
	for i := 1; i < n; i++ {
		if nums[i] == pre {
			cnt++
			if cnt > 2 {
				nums[i] = sign
			}
		} else {
			pre = nums[i]
			cnt = 1
		}
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != sign {
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
		}
		fast++
	}
	return slow
}
