package valid_parentheses

// https://leetcode-cn.com/problems/valid-parentheses/
import "strings"

type stack struct {
	arr []string
}

func (s *stack) push(in string) {
	s.arr = append(s.arr, in)
}

func (s *stack) pop() string {
	l := len(s.arr)
	if l == 0 {
		return ""
	}

	out := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return out
}

const (
	right = "]})"
	left  = "[{("
)

func isValid(s string) bool {
	ss := stack{[]string{}}
	for _, v := range s {
		idx := strings.Index(right, string(v))
		if idx >= 0 {
			get := ss.pop()
			if string(left[idx]) != get {
				return false
			}
		} else if strings.Contains(left, string(v)) {
			ss.push(string(v))
		} else {
			continue
		}
	}

	return len(ss.arr) <= 0
}
