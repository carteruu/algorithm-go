package leetcode

func countPrimes1(n int) int {
	arr := make([]bool, n)
	ans := 0
	for i := 2; i < n; i++ {
		if arr[i] {
			continue
		}
		ans++
		for j := 2 * i; j < n; j += i {
			arr[j] = true
		}
	}
	return ans
}

// 超时警告
func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes2(n int) (cnt int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}

func countPrimes3(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}

func countPrimes4(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}

func countPrimes(n int) int {
	//小于n，不是小于等于n
	if n < 3 {
		return 0
	}

	isComposite := make([]bool, n)
	count := n / 2

	for i := 3; i*i < n; i += 2 {
		if isComposite[i] {
			continue
		}

		for j := i * i; j < n; j += 2 * i {
			if !isComposite[j] {
				count--
				isComposite[j] = true
			}
		}
	}

	return count
}
