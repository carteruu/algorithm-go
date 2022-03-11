package leet

func countDigitOne(n int) int {
	d := 1
	ans := 0
	for n > 0 {
		ans = d + (n/10%10)*ans
		n /= 10
		d *= 10
	}
	return ans
}

func countDigitOne1(n int) (ans int) {
	// mulk 表示 10^k
	// 在下面的代码中，可以发现 k 并没有被直接使用到（都是使用 10^k）
	// 但为了让代码看起来更加直观，这里保留了 k
	for k, mulk := 0, 1; n >= mulk; k++ {
		ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
		mulk *= 10
	}
	return
}
