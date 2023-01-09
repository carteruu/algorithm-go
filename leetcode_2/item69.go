package leetcode_2

import "math"

func mySqrt(x int) int {
	l, r := 1, x
	for l <= r {
		mid := (l + r) >> 1
		if mid*mid < x {
			l = mid + 1
		} else if mid*mid > x {
			r = mid - 1
		} else {
			return mid
		}
	}
	return r
}

func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

//arr=[]
//a=document.querySelectorAll(".css-rxsf89");for(var i=0;i<a.length;i++){arr.push(a[i].innerHTML)}
//a=document.querySelectorAll(".css-i6szel");for(var i=0;i<a.length;i++){arr.push(a[i].innerHTML)}
//a=document.querySelectorAll(".css-5ikh7q");for(var i=0;i<a.length;i++){arr.push(a[i].innerHTML)}
//arr.sort((x,y)=>x-y)
