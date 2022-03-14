package leetcode

type PhoneDirectory struct {
	//使用状态
	uses map[int]struct{}
}

func Constructor(maxNumbers int) PhoneDirectory {
	//初始化使用状态,存在即可用,时间复杂度:O(n)
	uses := make(map[int]struct{}, maxNumbers)
	for i := 0; i < maxNumbers; i++ {
		uses[i] = struct{}{}
	}
	return PhoneDirectory{
		uses: uses,
	}
}

//获取,时间复杂度:O(1)
func (this *PhoneDirectory) Get() int {
	for i := range this.uses {
		//将i删除,并返回
		delete(this.uses, i)
		return i
	}
	return -1
}

//检查,时间复杂度:O(1)
func (this *PhoneDirectory) Check(number int) bool {
	_, ok := this.uses[number]
	return ok
}

//释放,时间复杂度:O(1)
func (this *PhoneDirectory) Release(number int) {
	this.uses[number] = struct{}{}
}
