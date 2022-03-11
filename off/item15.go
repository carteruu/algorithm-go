package off

func hammingWeight(num uint32) int {
	ans := 0
	for ; num > 0; num &= num - 1 {
		ans++
	}
	return ans
}
func hammingWeight2(num uint32) int {
	ans := 0
	for num > 0 {
		if num&1 == 1 {
			ans++
		}
		num >>= 1
	}
	return ans
}
func hammingWeight1(num uint32) int {
	ans := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			ans++
		}
		num >>= 1
	}
	return ans
}
