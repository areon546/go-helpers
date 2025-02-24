package gohelpers

import "strconv"

func ConvertToInteger(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, err
	}
	return i, err
}

// dont think they are necessary honestly but here still
func BytesToString(b []byte) (s string) {
	return string(b)
}
func StringToBytes(s string) (b []byte) {
	return []byte(s)
}
