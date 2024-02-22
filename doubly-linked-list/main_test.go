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
