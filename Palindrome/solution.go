package Palindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := 0
	for i := x; i > 0; {
		y = y*10 + i%10
		i = i / 10
	}
	return y == x
}
