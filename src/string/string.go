package main

func ReverseString(s string) string {
	ret := []rune(s)
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		// swap
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}