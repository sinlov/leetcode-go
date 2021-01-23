package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()
	ast.True(q.IsEmpty(), "NewQueue not empty")

	start, end := 0, 100

	for i := start; i < end; i++ {
		q.Push(i)
		ast.Equal(i-start+1, q.Len(), "Push queue length not equal")
	}

	for i := start; i < end; i++ {
		ast.Equal(i, q.Pop(), "pop from queue not equal")
	}

	ast.True(q.IsEmpty(), "final queue not empty")
}
