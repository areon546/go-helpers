package helpers

import "strconv"

/*
All functions in this class must conform to the nomeclature:

ObjectToConvertedObject

*/

func StringToInteger(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, err
	}
	return i, err
}

func BytesToString(b []byte) string {
	return string(b)
}
func StringToBytes(s string) []byte {
	return []byte(s)
}
