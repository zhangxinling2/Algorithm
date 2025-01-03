package main

func removeDuplicates(s string) string {
	bs := []byte(s)
	ms := make([]byte, 0)
	for _, b := range bs {
		if len(ms) == 0 {
			ms = append(ms, b)
			continue
		}
		if ms[len(ms)-1] == b {
			ms = ms[:len(ms)-1]
		} else {
			ms = append(ms, b)
		}
	}
	return string(ms)
}
