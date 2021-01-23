package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	ast := assert.New(t)

	s := NewStack()
	ast.True(s.IsEmpty(), "NewStack not empty")

	start, end := 0, 100

	for i := start; i < end; i++ {
		s.Push(i)
		ast.Equal(i-start+1, s.Len(), "Push stack length not right")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, s.Pop(), "Pop not equal")
	}

	ast.True(s.IsEmpty(), "final Pop not empty")
}
