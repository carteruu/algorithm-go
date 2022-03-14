package leetcode

type PhoneDirectory11 struct {
	//使用状态
	uses       []bool
	maxNumbers int
}

//初始化,使用状态默认为false:未使用,时间复杂度:O(1)
func ConstructorPhoneDirectory11(maxNumbers int) PhoneDirectory11 {
	return PhoneDirectory11{
		uses:       make([]bool, maxNumbers),
		maxNumbers: maxNumbers,
	}
}

//遍历,找到第一个未使用的号码.时间复杂度 :O(n)
func (this *PhoneDirectory11) Get() int {
	for i := 0; i < this.maxNumbers; i++ {
		if !this.uses[i] {
			this.uses[i] = true
			return i
		}
	}
	return -1
}

//判断是否被使用,时间复杂度:O(1)
func (this *PhoneDirectory11) Check(number int) bool {
	return !this.uses[number]
}

//释放,时间复杂度:O(1)
func (this *PhoneDirectory11) Release(number int) {
	this.uses[number] = false
}
