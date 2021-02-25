package _00071_simplify_path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock simplifyPath

var qs = []question{
	{
		param{"/..home"},
		answer{"/..home"},
	},
	{
		param{"/.home"},
		answer{"/.home"},
	},
	{
		param{"/a/./b/../../c/"},
		answer{"/c"},
	},
	{
		param{"/home//foo/"},
		answer{"/home/foo"},
	},
	{
		param{"/home/"},
		answer{"/home"},
	},
	{
		param{"/..."},
		answer{"/..."},
	},
	{
		param{"/../"},
		answer{"/"},
	},
}

func Test_simplifyPath(t *testing.T) {
	t.Logf("~> LeetCode simplifyPath start")
	// do simplifyPath
	for _, q := range qs {
		a, p := q.answer, q.param
		res := simplifyPath(p.one)
		// verify simplifyPath
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode simplifyPath end")
}

func Benchmark_simplifyPath(b *testing.B) {
	// do simplifyPath
	for _, q := range qs {
		simplifyPath(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one string
}

// one first  answer
type answer struct {
	one string
}
