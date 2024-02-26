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
	return preOrderWalk(root, []int{})
}

func preOrderWalk(cur *Node, path []int) []int {
	if cur == nil {
		return path
	}
	path = append(path, cur.val)

	path = preOrderWalk(cur.left, path)
	path = preOrderWalk(cur.right, path)

	return path
}

func inOrderSearch(root *Node) []int {
	return inOrderWalk(root, []int{})
}

func inOrderWalk(cur *Node, path []int) []int {
	if cur == nil {
		return path
	}

	path = inOrderWalk(cur.left, path)
	path = append(path, cur.val)
	path = inOrderWalk(cur.right, path)

	return path
}

func postOrderSearch(root *Node) []int {
	return postOrderWalk(root, []int{})
}

func postOrderWalk(cur *Node, path []int) []int {
	if cur == nil {
		return path
	}

	path = postOrderWalk(cur.left, path)
	path = postOrderWalk(cur.right, path)
	path = append(path, cur.val)

	return path
}
