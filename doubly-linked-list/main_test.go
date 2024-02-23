package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppend(t *testing.T) {
	c := require.New(t)

	l := NewDoublyLinkedList()
	err := l.Append(1)
	c.Nil(err)
	err = l.Append(2)
	c.Nil(err)
	err = l.Append(3)
	c.Nil(err)

	c.Equal("123", l.Print())
	c.Equal("321", l.ReversePrint())
}

func TestPrepend(t *testing.T) {
	c := require.New(t)

	l := NewDoublyLinkedList()
	err := l.Prepend(1)
	c.Nil(err)
	err = l.Prepend(2)
	c.Nil(err)
	err = l.Prepend(3)
	c.Nil(err)

	c.Equal("321", l.Print())
	c.Equal("123", l.ReversePrint())
}

func TestGet(t *testing.T) {
	c := require.New(t)

	l := NewDoublyLinkedList()
	err := l.Append(1)
	c.Nil(err)
	err = l.Append(2)
	c.Nil(err)
	err = l.Append(3)
	c.Nil(err)

	val, err := l.Get(1)
	c.Nil(err)
	c.Equal(2, val)

	val, err = l.Get(0)
	c.Nil(err)
	c.Equal(1, val)

	val, err = l.Get(2)
	c.Nil(err)
	c.Equal(3, val)

	_, err = l.Get(3)
	c.ErrorIs(err, errIndexTooLarge)
}

func TestInsertAt(t *testing.T) {
	c := require.New(t)

	l := NewDoublyLinkedList()
	err := l.Append(1)
	c.Nil(err)
	err = l.Append(2)
	c.Nil(err)
	err = l.Append(3)
	c.Nil(err)

	err = l.InsertAt(4, 1)
	c.Nil(err)
	c.Equal("1423", l.Print())
	c.Equal(4, l.length)

	err = l.InsertAt(0, 2)
	c.Nil(err)
	c.Equal("14023", l.Print())
	c.Equal(5, l.length)

	err = l.InsertAt(6, 0)
	c.Nil(err)
	c.Equal("614023", l.Print())
	c.Equal(6, l.length)

	err = l.InsertAt(7, 7)
	c.ErrorIs(err, errIndexTooLarge)
}

func TestRemoveAt(t *testing.T) {
	c := require.New(t)

	l := NewDoublyLinkedList()
	err := l.Append(1)
	c.Nil(err)
	err = l.Append(2)
	c.Nil(err)
	err = l.Append(3)
	c.Nil(err)

	item, err := l.RemoveAt(1)
	c.Nil(err)
	c.Equal(2, item)
	c.Equal("13", l.Print())
	c.Equal(2, l.length)

	item, err = l.RemoveAt(0)
	c.Nil(err)
	c.Equal(1, item)
	c.Equal("3", l.Print())
	c.Equal(1, l.length)

	item, err = l.RemoveAt(7)
	c.ErrorIs(err, errIndexTooLarge)

	item, err = l.RemoveAt(0)
	c.Nil(err)
	c.Equal(3, item)
	c.Equal("", l.Print())
	c.Equal(0, l.length)

	item, err = l.RemoveAt(0)
	c.ErrorIs(err, errListIsEmpty)
}
