package _00049_group_anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock groupAnagrams

var qs = []question{
	{
		param{[]string{
			"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc",
		}},
		answer{[][]string{
			{"cab"},
			{"duh"},
			{"ill"},
			{"bar"},
			{"doc"},
			{"tin"},
			{"pew"},
			{"may"},
			{"buy"},
			{"max"},
		}},
	},
	{
		param{[]string{
			"eat", "tea", "tan", "ate", "nat", "bat",
		}},
		answer{[][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}},
	},
}

func Test_groupAnagrams(t *testing.T) {
	t.Logf("~> LeetCode groupAnagrams start")
	// do groupAnagrams
	for _, q := range qs {
		a, p := q.answer, q.param
		res := groupAnagrams(p.one)
		// verify groupAnagrams
		assert.Equal(t, len(a.one), len(res),
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode groupAnagrams end")
}

func Benchmark_groupAnagrams(b *testing.B) {
	// do groupAnagrams
	for _, q := range qs {
		groupAnagrams(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []string
}

// one first  answer
type answer struct {
	one [][]string
}
