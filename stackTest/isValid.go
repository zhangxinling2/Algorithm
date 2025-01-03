package main

import "fmt"

func isValid() bool {
	var bs []byte
	fmt.Scanln(&bs)
	var ms MyStack[byte] = make([]byte, 0)
	m := make(map[byte]byte)
	m['}'] = '{'
	m[']'] = '['
	m[')'] = '('
	for i := 0; i < len(bs); i++ {
		if bs[i] == '(' || bs[i] == '{' || bs[i] == '[' {
			ms.Push(bs[i])
		} else {
			if len(ms) == 0 {
				return false
			}
			if ms.Peek() != m[bs[i]] {
				return false
			}
			ms.Pop()
		}

	}
	if len(ms) != 0 {
		return false
	}
	return true
}
