package main

import "fmt"

func stringRelay() {
	var (
		n    int
		str1 string
		str2 string
		tmp  string
	)
	fmt.Scanln(&n)
	fmt.Scanln(&str1, &str2)
	strMap := make(map[string]struct{}, 0)
	for i := 0; i < n; i++ {
		fmt.Scanln(&tmp)
		strMap[tmp] = struct{}{}
	}
	visitedMap := make(map[string]int, 0)
	visitedMap[str1] = 1
	queue := make([]string, 0)
	queue = append(queue, str1)
	for len(queue) > 0 {
		str := queue[0]
		queue = queue[1:]
		path := visitedMap[str]
		for i := 0; i < len(str); i++ {
			newWord := []byte(str)
			for j := 0; j < 26; j++ {
				newWord[i] = byte(j + 'a')
				if string(newWord) == str2 {
					fmt.Println(path + 1)
				}
				_, ok := strMap[string(newWord)]
				_, visited := visitedMap[string(newWord)]
				if ok && !visited {
					queue = append(queue, string(newWord))
					visitedMap[string(newWord)] = path + 1
				}
			}
		}
	}
}
