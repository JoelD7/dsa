package main

import (
	"errors"
	"strconv"
)

var (
	errIndexTooLarge = errors.New("index exceeds the size of the list")
	errListIsEmpty   = errors.New("list is empty")
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (d *DoublyLinkedList) Length() int {
	return d.length
}

func (d *DoublyLinkedList) InsertAt(item int, index int) error {
	if index == 0 {
		return d.Prepend(item)
	}

	if d.head == nil {
		return d.Append(item)
	}

	if index >= d.length {
		return errIndexTooLarge
	}

	d.length++

	cur := d.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}

	newItem := &Node{val: item, prev: cur.prev, next: cur}
	curPrev := cur.prev
	curPrev.next = newItem
	cur.prev = newItem

	return nil
}

// Remove removes the first element that matches item
func (d *DoublyLinkedList) Remove(item int) (int, error) {
	if d.head == nil {
		return 0, errListIsEmpty
	}

	d.length--

	cur := d.head
	for i := 0; i < d.length && cur != nil; i++ {
		if cur.val == item {
			break
		}

		cur = cur.next
	}

	curPrev := cur.prev
	curPrev.next = cur.next
	cur.next.prev = curPrev

	return cur.val, nil
}

func (d *DoublyLinkedList) RemoveAt(index int) (int, error) {
	if d.head == nil {
		return 0, errListIsEmpty
	}

	if index >= d.length {
		return 0, errIndexTooLarge
	}

	d.length--

	cur := d.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}

	if cur == d.head {
		d.head = cur.next
		return cur.val, nil
	}

	curPrev := cur.prev
	if curPrev != nil {
		curPrev.next = cur.next
	}

	cur.next.prev = curPrev

	return cur.val, nil
}

func (d *DoublyLinkedList) Append(item int) error {
	d.length++

	if d.head == nil {
		newItem := &Node{val: item}
		d.head = newItem
		d.tail = newItem
		return nil
	}

	newItem := &Node{item, d.tail, nil}
	d.tail.next = newItem
	d.tail = newItem

	return nil
}

func (d *DoublyLinkedList) Prepend(item int) error {
	d.length++

	if d.head == nil {
		newItem := &Node{val: item}
		d.head = newItem
		d.tail = newItem
		return nil
	}

	newItem := &Node{item, nil, d.head}
	d.head.prev = newItem
	d.head = newItem

	return nil
}

func (d *DoublyLinkedList) Get(index int) (int, error) {
	if d.head == nil {
		return 0, errListIsEmpty
	}

	if index >= d.length {
		return 0, errIndexTooLarge
	}

	cur := d.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}

	return cur.val, nil
}

func (d *DoublyLinkedList) Print() string {
	cur := d.head
	var s string

	for i := 0; i < d.length && cur != nil; i++ {
		s += strconv.FormatInt(int64(cur.val), 10)
		cur = cur.next
	}

	return s
}

func (d *DoublyLinkedList) ReversePrint() string {
	cur := d.tail
	var s string

	for i := 0; i < d.length && cur != nil; i++ {
		s += strconv.FormatInt(int64(cur.val), 10)
		cur = cur.prev
	}

	return s
}

func main() {
}
