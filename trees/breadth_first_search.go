package trees

type Queue struct {
	head   *QNode
	tail   *QNode
	length int
}

type QNode struct {
	next *QNode
	val  *Node
}

func BFS(root *Node) []int {
	queue := NewQueue()

	path := make([]int, 0)
	cur := root

	for {
		path = append(path, cur.val)

		if cur.left != nil {
			queue.Enqueue(cur.left)
		}

		if cur.right != nil {
			queue.Enqueue(cur.right)
		}

		cur = queue.Dequeue()

		if cur == nil {
			break
		}
	}

	return path
}

func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) Enqueue(val *Node) {
	q.length++

	if q.head == nil {
		newNode := &QNode{nil, val}
		q.head = newNode
		q.tail = newNode

		return
	}

	newNode := &QNode{nil, val}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) Dequeue() *Node {
	if q.head == nil {
		return nil
	}

	q.length--

	val := q.head.val
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return val
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}
