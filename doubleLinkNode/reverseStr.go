package main

func reverseStr(s string, k int) string {
	ss := []byte(s)
	l := len(ss)
	remainder := 0
	for i := 2 * k; i <= l; i = i + 2*k {
		reverse(ss[i-2*k : i-1])
		remainder = i
	}
	if l-remainder+1 < k {
		reverse(ss[remainder:])
	} else {
		reverse(ss[remainder : remainder+k])
	}
	return string(ss)
}
func reverse(str []byte) {
	r := len(str) / 2
	l := 0
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
}
