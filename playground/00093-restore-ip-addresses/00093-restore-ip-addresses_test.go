package _00093_restore_ip_addresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock restoreIpAddresses

var qs = []question{
	{
		param{"25525511135"},
		answer{[]string{"255.255.111.35", "255.255.11.135"}},
	},
	{
		param{"0000"},
		answer{[]string{"0.0.0.0"}},
	},
	{
		param{"1111"},
		answer{[]string{"1.1.1.1"}},
	},
	{
		param{"010010"},
		answer{[]string{"0.100.1.0", "0.10.0.10"}},
	},
	{
		param{"101023"},
		answer{[]string{"101.0.2.3", "10.10.2.3", "10.1.0.23", "1.0.102.3", "1.0.10.23"}},
	},
	{
		param{""},
		answer{[]string{}},
	},
}

func Test_restoreIpAddresses(t *testing.T) {
	t.Logf("~> LeetCode restoreIpAddresses start")
	// do restoreIpAddresses
	for _, q := range qs {
		a, p := q.answer, q.param
		res := restoreIpAddresses(p.one)
		// verify restoreIpAddresses
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode restoreIpAddresses end")
}

func Benchmark_restoreIpAddresses(b *testing.B) {
	// do restoreIpAddresses
	for _, q := range qs {
		restoreIpAddresses(q.param.one)
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
	one []string
}
