package leetcode

import "math"

func myAtoi(s string) int {
	a := NewAutomaton()
	for _, c := range s {
		if a.update(c) {
			break
		}
	}
	return int(a.ans())
}

type Automaton struct {
	sign  int64
	num   int64
	state int
	table [][]int
}

func NewAutomaton() *Automaton {
	return &Automaton{
		sign:  1,
		num:   0,
		state: 0,
		table: [][]int{
			{0, 1, 2, 3},
			{3, 3, 2, 3},
			{3, 3, 2, 3},
			{3, 3, 3, 3},
		},
	}
}
func (a *Automaton) update(c rune) bool {
	a.state = a.table[a.state][a.getNext(c)]
	if a.state == 1 && c == '-' {
		a.sign = -1
	} else if a.state == 2 {
		a.num = a.num*10 + int64(c-'0')
		if a.ans() > math.MaxInt32 {
			a.num = math.MaxInt32
			a.state = 3
		}
		if a.ans() < math.MinInt32 {
			a.num = -math.MinInt32
			a.state = 3
		}
	}
	return a.state == 3
}

func (a *Automaton) ans() int64 {
	return a.num * a.sign
}
func (a *Automaton) getNext(c rune) int {
	if c == ' ' {
		return 0
	} else if c == '+' || c == '-' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	} else {
		return 3
	}
}
