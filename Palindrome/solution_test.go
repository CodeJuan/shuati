package Palindrome

import "testing"


type ucTest struct {
	in       int
	expected bool
}

var ucTests = []ucTest{
	ucTest{1, true},
	ucTest{12, false},
	ucTest{121, true},
	ucTest{-1, false},
	ucTest{-121, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, ut := range ucTests {
		uc := isPalindrome(ut.in)
		if uc != ut.expected{
			t.Errorf("isPalindrome(%v) = %v,must be %v", ut.in, uc, ut.expected)
		}
	}
}