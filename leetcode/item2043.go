package leetcode

type Bank struct {
	balance []int64
	n       int
}

func ConstructorBank(balance []int64) Bank {
	return Bank{
		balance: balance,
		n:       len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.checkAccount(account1) || !this.checkAccount(account2) {
		return false
	}
	if !this.checkMoney(account1, money) {
		return false
	}
	this.Withdraw(account1, money)
	this.Deposit(account2, money)
	return true
}

func (this *Bank) checkMoney(account1 int, money int64) bool {
	if this.balance[this.getIdx(account1)] < money {
		return false
	}
	return true
}

func (this *Bank) checkAccount(account int) bool {
	if account < 1 || account > this.n {
		return false
	}
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.checkAccount(account) {
		return false
	}
	this.balance[this.getIdx(account)] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.checkAccount(account) {
		return false
	}
	if !this.checkMoney(account, money) {
		return false
	}
	this.balance[this.getIdx(account)] -= money
	return true
}

func (this *Bank) getIdx(account int) int {
	return account - 1
}
