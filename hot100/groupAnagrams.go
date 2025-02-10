package main

import (
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		t := []byte(str)
		slices.Sort(t)
		sortedS := string(t)
		m[sortedS] = append(m[sortedS], str)
	}
	return slices.Collect(maps.Values(m))
}
