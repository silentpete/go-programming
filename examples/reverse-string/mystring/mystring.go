package mystring

func Reverse(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		j := len(rs) - i - 1
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
