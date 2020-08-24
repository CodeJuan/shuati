package count

import "testing"

type ucTest struct {
	in       int
	expected string
}

var ucTests = []ucTest{
	ucTest{1, "1"},
	ucTest{2, "11"},
	ucTest{3, "21"},
	ucTest{4, "1211"},
	ucTest{5, "111221"},
	ucTest{6, "31221"},
	ucTest{7, "13112211"},
}

func TestCountAndSay(t *testing.T) {
	for _, ut := range ucTests {
		uc := countAndSay(ut.in)
		if uc != ut.expected {
			t.Errorf("%v = %v,must be %v", ut.in, uc, ut.expected)
		}
	}
}
