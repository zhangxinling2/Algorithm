package main

func removeExtraSpaces(s string) string {
	res := make([]byte, 0, len(s))
	pre, last := len(s)-1, len(s)-1
	for s[last] == ' ' {
		pre--
		last--
	}
	tmpByte := []byte(s)
	for pre >= 0 {
		last = pre
		for pre >= 0 && s[pre] >= 'a' && s[pre] <= 'z' {
			pre--
		}
		res = append(res, tmpByte[pre+1:last+1]...)
		for pre >= 0 && s[pre] == ' ' {
			pre--
		}
		res = append(res, ' ')
	}
	return string(res)
}
