package main

func letterCombinations(digits string) []string {
	res := []string{}
	var rb []byte
	ltd := make(map[byte]string)
	ltd['2'] = "abc"
	ltd['3'] = "def"
	ltd['4'] = "ghi"
	ltd['5'] = "jkl"
	ltd['6'] = "mno"
	ltd['7'] = "pqrs"
	ltd['8'] = "tuv"
	ltd['9'] = "wxyz"
	bs := []byte(digits)
	var backtracking func(bs []byte, index int)
	backtracking = func(bs []byte, index int) {
		if len(rb) == len(bs) {
			res = append(res, string(rb))
			return
		}
		tmp := []byte(ltd[bs[index]])
		for j := 0; j < len(tmp); j++ {
			rb = append(rb, tmp[j])
			backtracking(bs, index+1)
			rb = rb[:len(rb)-1]
		}
	}
	backtracking(bs, 0)
	return res
}
