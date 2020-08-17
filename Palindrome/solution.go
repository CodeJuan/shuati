package Palindrome

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	s2 := ""
	for _, v := range s{
		s2 = string(v) + s2
	}
	if strings.Compare(s ,s2) == 0 {
		return true
	}
	return false
}