package helpers

func ReverseString(s string) string {
	rs := ""

	for i := len(s) - 1; i > -1; i-- {
		rs += string(s[i])
	}

	return rs
}
