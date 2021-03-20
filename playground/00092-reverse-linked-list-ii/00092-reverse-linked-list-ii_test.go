package _00092_reverse_linked_list_ii

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock reverseBetween

var qs = []question{
	{
		param{
			one:   []int{1, 2, 3, 4, 5},
			two:   1,
			three: 2,
		},
		answer{[]int{2, 1, 3, 4, 5}},
	},
	{
		param{
			one:   []int{1, 2, 3, 4, 5},
			two:   2,
			three: 4,
		},
		answer{[]int{1, 4, 3, 2, 5}},
	},
	{
		param{
			one:   []int{5},
			two:   1,
			three: 1,
		},
		answer{[]int{5}},
	},
	{
		param{
			one:   []int{},
			two:   1,
			three: 1,
		},
		answer{[]int{}},
	},
}

func Test_reverseBetween(t *testing.T) {
	t.Logf("~> LeetCode reverseBetween start")
	// do reverseBetween
	for _, q := range qs {
		a, p := q.answer, q.param
		minLN := structures.IntList2ListNode(p.one)
		resLN := reverseBetween(minLN, p.two, p.three)
		res := structures.List2IntList(resLN)
		// verify reverseBetween
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v, %v ], out [%v] , want [ %v ]", p.one, p.two, p.three, res, a.one)
	}
	t.Logf("~> LeetCode reverseBetween end")
}

func Benchmark_reverseBetween(b *testing.B) {
	// do reverseBetween
	for _, q := range qs {
		minLN := structures.IntList2ListNode(q.param.one)
		reverseBetween(minLN, q.param.two, q.param.three)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one   []int
	two   int
	three int
}

// one first  answer
type answer struct {
	one []int
}
