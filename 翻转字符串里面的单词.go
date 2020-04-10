package leetcode

func reverseWords(s string) string {
	found := false
	last := 0
	ans := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		if !found && s[i] == ' ' {
			continue
		}
		if s[i] != ' ' {
			if !found {
				found = true
				last = i
			}
			continue
		}
		if found && s[i] == ' ' {
			for j := i + 1; j <= last; j++ {
				ans = append(ans, s[j])
			}
			ans = append(ans, ' ')
			found = false
		}
	}
	if found {
		for j := 0; j <= last; j++ {
			ans = append(ans, s[j])
		}
	} else if len(ans) != 0 {
		ans = ans[:len(ans)-1]
	}
	return string(ans)
}
