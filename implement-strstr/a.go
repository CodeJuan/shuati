package implement

// https://leetcode-cn.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}

	idxHay := -1
	idxNeed := 0
	lenNeed := len(needle)
	i := 0
	for ; i < len(haystack); i++ {
		for j := idxNeed; j < lenNeed; j++ {
			if haystack[i] == needle[j] {
				if idxHay == -1 {
					idxHay = i
				}

				if j == lenNeed-1 {
					return idxHay
				}

				idxNeed = j + 1
				break
			} else {
				if idxHay >= 0 {
					i = idxHay
				}
				idxHay = -1
				idxNeed = 0
				break
			}
		}
	}

	return -1
}
