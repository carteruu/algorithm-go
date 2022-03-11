package off

import "strconv"

func translateNum(num int) int {
	dp1 := 1
	dp2 := 1
	for num > 0 {
		dp3 := dp2
		if num >= 10 {
			t := ((num/10)%10)*10 + (num % 10)
			if t >= 10 && t <= 25 {
				dp3 += dp1
			}
		}
		dp1, dp2 = dp2, dp3
		num /= 10
	}
	return dp2
}
func translateNum1(num int) int {
	dp1 := 1
	dp2 := 1
	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		dp3 := dp2
		if i < len(str)-1 {
			t := (str[i]-'0')*10 + (str[i+1] - '0')
			if t >= 10 && t <= 25 {
				dp3 += dp1
			}
		}
		dp1, dp2 = dp2, dp3
	}
	return dp2
}
