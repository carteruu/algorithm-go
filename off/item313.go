package off

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m := len(primes)
	pointers := make([]int, m)
	for i := range pointers {
		pointers[i] = 1
	}
	for i := 2; i <= n; i++ {
		nums := make([]int, m)
		minNum := math.MaxInt64
		for j, p := range pointers {
			nums[j] = dp[p] * primes[j]
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j, num := range nums {
			if minNum == num {
				pointers[j]++
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nthSuperUglyNumber1(n int, primes []int) int {
	//质数
	ps := make(map[int]struct{})
	noPrime := make(map[int]struct{})
	max := math.MaxInt32
	var isPrime func(a int) bool
	isPrime = func(a int) bool {
		if a < 2 || a >= max {
			return false
		}
		if _, ok := ps[a]; ok {
			return true
		}
		if _, ok := noPrime[a]; ok {
			return false
		}
		for i := 2; i*i <= a; i++ {
			if a%i == 0 {
				noPrime[a] = struct{}{}
				return false
			}
		}
		ps[a] = struct{}{}
		return true
	}

	primeMap := make(map[int]struct{}, len(primes))
	for _, prime := range primes {
		primeMap[prime] = struct{}{}
	}

	var res int
	p := 0
	for i := 1; p < n; i++ {
		is := true
		//找出所有质因数
		for j := 1; j*j <= i; j++ {
			if i%j != 0 {
				//不是因子
				continue
			}
			if _, ok := primeMap[j]; !ok && isPrime(j) {
				is = false
				break
			}
			if _, ok := primeMap[i/j]; !ok && isPrime(i/j) {
				is = false
				break
			}
		}
		if is {
			p++
			res = i
		}
	}
	return res
}
