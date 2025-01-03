package main

import "fmt"

func replaceNumber() string {
	var strByte []byte
	fmt.Scanln(&strByte)
	oldSize := len(strByte)
	for _, v := range strByte {
		if v > '0' && v < '9' {
			strByte = append(strByte, []byte("     ")...)
		}
	}
	tmpBytes := []byte("number")
	leftP, rightP := oldSize-1, len(strByte)-1
	for leftP < rightP {
		rightShift := 1
		if strByte[leftP] >= '0' && strByte[leftP] <= '9' {
			for i, tmpByte := range tmpBytes {
				strByte[rightP-len(tmpBytes)+i+1] = tmpByte
			}
			rightShift = len(tmpBytes)
		} else {
			strByte[rightP] = strByte[leftP]
		}
		// 更新指针
		rightP -= rightShift
		leftP -= 1
	}
	return string(strByte)
}
