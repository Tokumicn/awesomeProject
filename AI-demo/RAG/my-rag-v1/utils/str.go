package utils

// RuneCountInString 计算字符串的长度
func RuneCountInString(s string) (n int) {
	for range s {
		n++
	}
	return n
}
