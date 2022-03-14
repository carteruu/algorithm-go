package practise1

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		if ok1 != ok2 || val1 != val2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	t1 := &Tree{
		Value: 3,
		Left: &Tree{
			Value: 3,
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{}},
				Right: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{},
				},
			},
			Right: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{}},
				Right: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{},
				},
			}},
		Right: &Tree{
			Value: 3,
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{}},
				Right: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{},
				},
			},
			Right: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{}},
				Right: &Tree{
					Value: 3,
					Left:  &Tree{},
					Right: &Tree{},
				},
			},
		},
	}
	fmt.Print(Same(t1, t1))
}
