package longest_common_prefix

import (
	"strings"
	"testing"
)

type ucTest struct {
	in       []string
	expected string
}

var ucTests = []ucTest{
	ucTest{[]string{"fla", "flb", "flcd"}, "fl"},
	ucTest{[]string{"a", "b"}, ""},
	ucTest{[]string{"", "b"}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, ut := range ucTests {
		uc := longestCommonPrefix(ut.in)
		if strings.Compare(uc, ut.expected) != 0 {
			t.Errorf("isUnique(%v) = %v,must be %v", ut.in, uc, ut.expected)
		}
	}
}
