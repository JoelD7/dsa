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
