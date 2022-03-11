package leet

import "testing"

func TestBank_Transfer(t *testing.T) {
	type fields struct {
		balance []int64
		n       int
	}
	type args struct {
		account1 int
		account2 int
		money    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "0",
			fields: fields{
				balance: []int64{10, 100, 20, 50, 30},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := []int64{10, 100, 20, 50, 30}
			this := ConstructorBank(b)
			res := this.Transfer(1, 2, 20)
			if res != false {
				t.Errorf("Transfer() = %v, want %v", res, true)
			}

			res1 := this.Deposit(1, 10)
			if res1 != true {
				t.Errorf("Deposit() = %v, want %v", res, true)
			}

			res2 := this.Transfer(1, 2, 20)
			if res2 != true {
				t.Errorf("Transfer() = %v, want %v", res, true)
			}

			res3 := this.Withdraw(1, 10)
			if res3 != false {
				t.Errorf("Deposit() = %v, want %v", res, false)
			}

			if got := this.Transfer(tt.args.account1, tt.args.account2, tt.args.money); got != tt.want {
				t.Errorf("Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
