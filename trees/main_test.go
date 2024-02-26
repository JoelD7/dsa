package trees

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPreOrder(t *testing.T) {
	c := require.New(t)

	root := getTestTree()

	path := preOrderSearch(root)
	c.Equal([]int{7, 23, 5, 4, 3, 18, 21}, path)
}

func TestInOrder(t *testing.T) {
	c := require.New(t)

	root := getTestTree()

	path := inOrderSearch(root)
	c.Equal([]int{5, 23, 4, 7, 18, 3, 21}, path)
}

func TestPostOrder(t *testing.T) {
	c := require.New(t)

	root := getTestTree()

	path := postOrderSearch(root)
	c.Equal([]int{5, 4, 23, 18, 21, 3, 7}, path)
}

func getTestTree() *Node {
	root := NewNode(7)
	root.left = &Node{
		val: 23,
		left: &Node{
			5,
			nil,
			nil,
		},
		right: &Node{
			4,
			nil,
			nil,
		},
	}
	root.right = &Node{
		val: 3,
		left: &Node{
			18,
			nil,
			nil,
		},
		right: &Node{
			21,
			nil,
			nil,
		},
	}

	return root
}
