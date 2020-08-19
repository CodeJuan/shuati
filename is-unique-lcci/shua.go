package is_unique_lcci

const (
	CHAR_COUNT = 26
)

func isUnique(astr string) bool {
	a := make([]int, CHAR_COUNT*2)
	for _, v := range astr {
		if v >= 'a' && v <= 'z' {
			a[v-'a'] = a[v-'a'] + 1
		}
		if v >= 'A' && v <= 'Z' {
			a[v-'A'+CHAR_COUNT] = a[v-'a'+CHAR_COUNT] + 1
		}
	}
	for _, v := range a {
		if v > 1 {
			return false
		}
	}

	return true
}
