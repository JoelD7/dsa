package trees

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(root int) *Node {
	return &Node{
		val:   root,
		left:  nil,
		right: nil,
	}
}

func preOrderSearch(root *Node) []int {
	return walk(root, []int{})
}

func walk(cur *Node, path []int) []int {
	if cur == nil {
		return path
	}
	path = append(path, cur.val)

	path = walk(cur.left, path)
	path = walk(cur.right, path)

	return path
}
