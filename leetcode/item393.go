package leetcode

func validUtf8(data []int) bool {
	start := true
	cnt10 := 0
	for _, d := range data {
		if start {
			if d&0b10000000 == 0 {
				continue
			}
			start = false
			if d&0b11111000 == 0b11110000 {
				cnt10 = 3
			} else if d&0b11110000 == 0b11100000 {
				cnt10 = 2
			} else if d&0b11100000 == 0b11000000 {
				cnt10 = 1
			} else {
				return false
			}
		} else {
			if d&0b11000000 == 0b10000000 && cnt10 > 0 {
				cnt10--
				if cnt10 == 0 {
					start = true
				}
			} else {
				return false
			}
		}
	}
	return start
}
