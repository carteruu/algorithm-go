package leet

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(cur int) int {
		return ((cur+nums[cur])%n + n) % n
	}

	for i, num := range nums {
		if num == 0 {
			continue
		}
		slow, fast := i, next(i)
		// 判断非零且方向相同
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				if slow == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		add := i
		for nums[add]*nums[next(add)] > 0 {
			tmp := add
			add = next(add)
			nums[tmp] = 0
		}
	}
	return false
}

func circularArrayLoop1(nums []int) bool {
	n := len(nums)
	var hasLoop func(start int) bool
	hasLoop = func(start int) bool {
		var nextPos func(pos int) int
		nextPos = func(t int) int {
			//往回走过头
			for t < 0 {
				t += n
			}
			//往后走过尾
			for t >= n {
				t -= n
			}
			return t
		}

		times := 0
		slow := start
		fast := start
		for {
			slow += nums[slow]
			slow = nextPos(slow)

			fast += nums[fast]
			fast = nextPos(fast)
			if slow == fast {
				slow, fast = fast, slow
				times++
				if times == 2 {
					break
				}
			}
			fast += nums[fast]
			fast = nextPos(fast)
			if slow == fast {
				slow, fast = fast, slow
				times++
				if times == 2 {
					break
				}
			}
		}

		cur := nextPos(slow + nums[slow])
		l := 1
		for cur != slow {
			if l > 0 {
				if nums[cur] < 0 {
					return false
				}
			}
			if l < 0 {
				if nums[cur] > 0 {
					return false
				}
			}
			cur = nextPos(cur + nums[cur])
			l++
		}
		return l/2 > 1
	}

	for i := 0; i < len(nums); i++ {
		if hasLoop(i) {
			return true
		}
	}
	return false
}
