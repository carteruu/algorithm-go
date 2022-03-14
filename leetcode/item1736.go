package leetcode

import "strconv"

func maximumTime(time string) string {
	var str string
	s1 := time[:1]
	s2 := time[1:2]
	s3 := time[3:4]
	s4 := time[4:5]
	//小时
	if s1 == "?" && s2 == "?" {
		str += "23"
	} else if s1 == "?" {
		if num, _ := strconv.Atoi(s2); num > 3 {
			str += "1"
		} else {
			str += "2"
		}
		str += s2
	} else if s2 == "?" {
		str += s1
		if num, _ := strconv.Atoi(s1); num > 1 {
			str += "3"
		} else {
			str += "9"
		}
	} else {
		str += s1 + s2
	}
	str += ":"
	//分钟
	if s3 == "?" {
		str += "5"
	} else {
		str += s3
	}
	if s4 == "?" {
		str += "9"
	} else {
		str += s4
	}
	return str
}
