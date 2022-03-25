package leetcode_1

import "sort"

func threeSum_1(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	sort.Ints(nums)
	//第一个数必然小于等于0
	for i := 0; i < n-2 && nums[i] <= 0; i++ {
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//双指针
		j, k := i+1, n-1
		//第三个数必须大于等于0
		for j < k && nums[k] >= 0 {
			//当第一个数确定，只需要处理第二个数的去重，则第三个数必然不重复
			//这里注意第二个数的边界，j 必须小于 k
			for j > i+1 && j < k && nums[j] == nums[j-1] {
				j++
			}
			//第三个数还有没有得选
			if j >= k {
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ans
}
