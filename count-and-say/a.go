package count

import "fmt"

// https://leetcode-cn.com/problems/count-and-say/

func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	out := ""
	for i := 1; i <= n; i++ {
		if i == 1 {
			out = "1"
			continue
		}
		prev := ""
		tmp := ""
		count := 0
		for j, v := range out {
			if j == 0 {
				prev = string(v)
				count = 1
				continue
			}
			if string(v) != prev {
				tmp = tmp + fmt.Sprintf("%d%s", count, prev)
				count = 1
				prev = string(v)
			} else {
				count++
			}
		}
		tmp = tmp + fmt.Sprintf("%d%s", count, prev)
		out = tmp
	}
	return out
}
