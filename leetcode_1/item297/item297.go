package leetcode_1

import (
	"bytes"
	"strconv"
	"strings"
)

const seq = "|"

var nilNode = &TreeNode{}

const nilNodeVal = "nil"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	q := []*TreeNode{root}
	bs := bytes.Buffer{}
	for len(q) > 0 {
		size := len(q)
		allNil := true
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			val := nilNodeVal
			if poll != nilNode {
				val = strconv.Itoa(poll.Val)
			}
			bs.WriteString(val)
			bs.WriteString(seq)
			left := poll.Left
			right := poll.Right
			if left != nil || right != nil {
				allNil = false
			}
			if left == nil {
				left = nilNode
			}
			if right == nil {
				right = nilNode
			}
			q = append(q, left, right)
		}
		if allNil {
			break
		}
	}
	rt := bs.String()
	return rt[:len(rt)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	return this.des(strings.Split(data, seq), 0)
}

func (this Codec) des(strs []string, idx int) *TreeNode {
	if idx >= len(strs) || strs[idx] == nilNodeVal {
		return nil
	}
	val, _ := strconv.Atoi(strs[idx])
	node := &TreeNode{Val: val}
	node.Left = this.des(strs, idx*2+1)
	node.Right = this.des(strs, idx*2+2)
	return node
}
