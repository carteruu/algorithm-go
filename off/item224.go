package off

func calculate(s string) int {
	ans := 0
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}
func calculate1(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	var cal func(start int) (int, int)
	cal = func(start int) (int, int) {
		isPositive := true
		num := 0
		res := 0
		var i int
		for i = start; i < n; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				num *= 10
				num += int(s[i] - '0')
			} else {
				if !isPositive {
					num = -num
				}
				res += num
				num = 0

				switch s[i] {
				case ' ':
					{
						continue
					}
				case '-':
					{
						isPositive = false
					}
				case '+':
					{
						isPositive = true
					}
				case '(':
					{
						rs, end := cal(i + 1)
						if !isPositive {
							rs = -rs
						}
						res += rs
						i = end
					}
				case ')':
					{
						if !isPositive {
							num = -num
						}
						res += num
						return res, i
					}
				}
			}
		}
		if !isPositive {
			num = -num
		}
		res += num
		return res, i
	}
	res, _ := cal(0)
	return res
}
