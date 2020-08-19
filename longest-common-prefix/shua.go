package longest_common_prefix

import "strings"

func longestCommonPrefix(strs []string) string {
	minLen := 0
	returnStr := ""
	for i, str := range strs {
		len := len(str)
		if len == 0 {
			return ""
		}
		if i == 0 {
			minLen = len
		} else {
			if len < minLen {
				minLen = len
			}
		}
	}
	for i := 0; i < minLen; i++ {
		var currentChar string
		for j, str := range strs {
			if j == 0 {
				currentChar = string(str[i])
			} else {
				if strings.Compare(string(str[i]), currentChar) != 0 {
					return returnStr
				}
			}
		}
		returnStr = returnStr + currentChar
	}
	return returnStr
}
