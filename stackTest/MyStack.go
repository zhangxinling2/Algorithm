package main

type MyStack[T int | byte] []T

func (s *MyStack[T]) Push(val T) {
	*s = append(*s, val)
}
func (s *MyStack[T]) Pop() T {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}
func (s *MyStack[T]) Peek() T {

	return (*s)[len(*s)-1]
}
