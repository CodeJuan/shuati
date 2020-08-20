package valid_parentheses

import "testing"

type ucTest struct {
	in       string
	expected bool
}

var ucTests = []ucTest{
	ucTest{"", true},
	ucTest{"{}", true},
	ucTest{"()[]{}", true},
	ucTest{"(]", false},
	ucTest{"(", false},
	ucTest{"]", false},
	ucTest{"{[]}", true},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, ut := range ucTests {
		uc := isValid(ut.in)
		if uc != ut.expected {
			t.Errorf("%v = %v,must be %v", ut.in, uc, ut.expected)
		}
	}
}
