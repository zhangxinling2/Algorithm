package main

import "fmt"

func rightVerse() string {
	var k int
	var str []byte
	fmt.Scanln(&str)
	fmt.Scan(&k)
	reverse(str)
	reverse(str[:k])
	reverse(str[k:])
	return string(str)
}
func reverse(str []byte) {
	j := len(str) - 1
	for i := 0; i < j; i++ {
		str[i], str[j] = str[j], str[i]
		j--
	}
}
