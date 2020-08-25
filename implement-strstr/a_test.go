package implement

import "testing"

type ucTest struct {
	haystack string
	needle   string
	expected int
}

var ucTests = []ucTest{
	ucTest{"mississippi", "pi", 9},
	ucTest{"mississippi", "issip", 4},
	ucTest{"hello", "ll", 2},
	ucTest{"hlello", "ll", 3},
	ucTest{"1", "2", -1},
	ucTest{"1", "", 0},
	ucTest{"", "1", -1},
	ucTest{"aba", "ac", -1},
	ucTest{"lhehelao", "la", 5},
	ucTest{"lhehelao", "lh", 0},
	ucTest{"lhelh", "lh", 0},
}

func TestCountAndSay(t *testing.T) {
	for _, ut := range ucTests {
		uc := strStr(ut.haystack, ut.needle)
		if uc != ut.expected {
			t.Errorf("%v, %v = %v,must be %v", ut.haystack, ut.needle, uc, ut.expected)
		}
	}
}
