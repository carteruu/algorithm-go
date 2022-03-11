package delink

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}
