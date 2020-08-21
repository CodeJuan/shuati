package remove

import "testing"

type ucTest struct {
	in       []int
	expected int
}

var ucTests = []ucTest{
	ucTest{[]int{}, 0},
	ucTest{[]int{1}, 1},
	ucTest{[]int{1, 2}, 2},
	ucTest{[]int{1, 2, 2, 3}, 3},
	ucTest{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, ut := range ucTests {
		uc := removeDuplicates(ut.in)
		if uc != ut.expected {
			t.Errorf("%v = %v,must be %v", ut.in, uc, ut.expected)
		}
	}
}
