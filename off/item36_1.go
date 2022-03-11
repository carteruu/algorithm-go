package off

type Node struct {
	Val         int
	Left, Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	var dfs func(node *Node) (head, tail *Node)
	dfs = func(node *Node) (head, tail *Node) {
		//默认
		head, tail = node, node
		if node.Left != nil {
			//递归,左子树
			lHead, lTail := dfs(node.Left)
			//如果左子树不为空,头结点为左子树的头节点.并拼接当前节点到左子树的尾节点
			head = lHead
			lTail.Right = node
			node.Left = lTail
		}
		if node.Right != nil {
			//递归,右子树
			rHead, rTail := dfs(node.Right)
			//如果右子树不为空,尾节点为右子树的尾节点.并拼接当前节点到右子树的头结点
			tail = rTail
			node.Right = rHead
			rHead.Left = node
		}
		return
	}
	head, tail := dfs(root)
	//最后将返回的头尾节点拼接成环
	tail.Right = head
	head.Left = tail
	return head
}
