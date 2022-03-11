package leet

import "bytes"

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	ans := make([]byte, n)
	idx := 0
	cycle := numRows*2 - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycle {
			ans[idx] = s[j+i]
			idx++
			//j+i+cycle-2*i或j + i + 2*(numRows-1-i)
			next := j + i + 2*(numRows-1-i)
			if i != 0 && i != numRows-1 && next < n {
				ans[idx] = s[next]
				idx++
			}
		}
	}
	return string(ans)
}
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	cycle := numRows + numRows - 2
	//生成每一行的Buffer
	bs := make([]*bytes.Buffer, cycle)
	//竖直向下方向
	for i := 0; i < numRows && i < len(s); i++ {
		bs[i] = &bytes.Buffer{}
	}
	//斜向上方向
	for i := 0; i < numRows-2; i++ {
		bs[numRows+i] = bs[numRows-i-2]
	}
	//按字符遍历
	cycleIdx := 0
	for _, c := range s {
		bs[cycleIdx].WriteRune(c)
		cycleIdx++
		cycleIdx %= cycle
	}
	var ans bytes.Buffer
	for i := 0; i < numRows && bs[i] != nil; i++ {
		ans.Write(bs[i].Bytes())
	}
	return ans.String()
}
