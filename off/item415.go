package off

import "strconv"

func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

func addStrings1(num1 string, num2 string) string {
	end1 := len(num1) - 1
	end2 := len(num2) - 1

	res := []rune{}
	var p rune = 0
	for end1 >= 0 || end2 >= 0 {
		c1 := '0'
		if end1 >= 0 {
			c1 = rune(num1[end1])
		}
		c2 := '0'
		if end2 >= 0 {
			c2 = rune(num2[end2])
		}
		c := c1 + (c2 - '0') + p
		if c > '9' {
			p = 1
			c -= 10
		} else {
			p = 0
		}
		res = append(res, c)
		end1--
		end2--
	}
	if p == 1 {
		res = append(res, '1')
	}
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}
