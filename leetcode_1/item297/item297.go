package leetcode_1

import (
	"bytes"
	"strconv"
	"strings"
)

const seq = "|"

const nilNodeVal = "nil"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	bs := &bytes.Buffer{}
	this.se(root, bs)
	return bs.String()
}
func (this *Codec) se(node *TreeNode, bs *bytes.Buffer) {
	if node == nil {
		bs.WriteString(nilNodeVal)
		bs.WriteString(seq)
		return
	}
	bs.WriteString(strconv.Itoa(node.Val))
	bs.WriteString(seq)
	this.se(node.Left, bs)
	this.se(node.Right, bs)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	root, _ := this.de(strings.Split(data, seq))
	return root
}

func (this Codec) de(strs []string) (*TreeNode, []string) {
	if len(strs) == 0 {
		return nil, nil
	}
	str := strs[0]
	strs = strs[1:]
	if str == nilNodeVal {
		return nil, strs
	}
	val, _ := strconv.Atoi(str)
	node := &TreeNode{
		Val: val,
	}
	node.Left, strs = this.de(strs)
	node.Right, strs = this.de(strs)
	return node, strs
}
