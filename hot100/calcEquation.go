package main

import "container/list"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	letters := map[string]*list.List{}
	visited := make(map[string]bool)
	for i, v := range equations {
		visited[v[0]] = false
		if val, ok := letters[v[0]]; ok {
			val.PushBack([]any{v[1], values[i]})
		} else {
			l := list.New()
			l.PushBack([]any{v[1], values[i]})
			letters[v[0]] = l
		}
		if val, ok := letters[v[1]]; ok {
			val.PushBack([]any{v[0], 1 / values[i]})
		} else {
			l := list.New()
			l.PushBack([]any{v[0], 1 / values[i]})
			letters[v[1]] = l
		}
	}
	res := make([]float64, 0, len(queries))

	var dfs func(resVal float64, index string, targetStr string) bool
	dfs = func(resVal float64, index string, targetStr string) bool {
		l := letters[index].Front()
		for l != nil {
			var isStop bool
			val, divisor := l.Value.([]any)[1].(float64), l.Value.([]any)[0].(string)
			if !visited[divisor] {
				visited[divisor] = true
				if divisor == targetStr {
					res = append(res, resVal*val)
					visited[divisor] = false
					return true
				} else {
					isStop = dfs(resVal*val, divisor, targetStr)
				}
				visited[divisor] = false
				if isStop {
					return true
				}
			}
			l = l.Next()
		}
		return false
	}

	for _, v := range queries {
		dividend, divisor := v[0], v[1]
		_, ok := letters[dividend]
		_, ok2 := letters[divisor]
		if !ok || !ok2 {
			res = append(res, -1.0)
			continue
		}
		if dividend == divisor {
			res = append(res, 1.0)
			continue
		}
		visited[dividend] = true
		t := dfs(1, dividend, divisor)
		if !t {
			res = append(res, -1.0)
		}
		visited[dividend] = false
	}
	return res
}
