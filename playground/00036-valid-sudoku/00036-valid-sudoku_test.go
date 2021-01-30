package _00036_valid_sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock isValidSudoku

var qs = []question{
	{
		param{
			[][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
		answer{false},
	},
	{
		param{
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
		answer{true},
	},
	{
		param{
			[][]byte{
				{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
				{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
		},
		answer{true},
	},
}

func Test_isValidSudoku(t *testing.T) {
	t.Logf("~> LeetCode isValidSudoku start")
	// do isValidSudoku
	for _, q := range qs {
		a, p := q.answer, q.param
		res := isValidSudoku(p.one)
		// verify isValidSudoku
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode isValidSudoku end")
}

func Benchmark_isValidSudoku(b *testing.B) {
	// do isValidSudoku
	for _, q := range qs {
		isValidSudoku(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one [][]byte
}

// one first  answer
type answer struct {
	one bool
}
