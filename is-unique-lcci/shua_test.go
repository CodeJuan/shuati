package is_unique_lcci

import "testing"

type ucTest struct {
	in       string
	expected bool
}

var ucTests = []ucTest{
	ucTest{"a", true},
	ucTest{"ab", true},
	ucTest{"aab", false},
	ucTest{"", true},
	ucTest{"abb", false},
}

func TestIsUnique(t *testing.T) {
	for _, ut := range ucTests {
		uc := isUnique(ut.in)
		if uc != ut.expected {
			t.Errorf("isUnique(%v) = %v,must be %v", ut.in, uc, ut.expected)
		}
	}
}
