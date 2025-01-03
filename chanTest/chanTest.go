package main

import (
	"fmt"
	"time"
)

type student struct {
	name string
}

func printInt(u <-chan int) {
	time.Sleep(time.Second * 2)
	fmt.Println("printPeople", <-u)
}
func printStu(u <-chan *student) {
	time.Sleep(time.Second * 2)
	fmt.Println("printPeople", <-u)
}

var stu = student{name: "A"}

func main() {
	//c := make(chan *student, 5)
	//a := &stu
	//c <- a
	//fmt.Println(a)
	//a = &student{name: "B"}
	//go func() {
	//	printStu(c)
	//}()
	//time.Sleep(time.Second * 3)
	//fmt.Println(a)
	c := make(chan int, 5)
	a := 1
	c <- a
	fmt.Println(a)
	a = 2
	go func() {
		printInt(c)
	}()
	time.Sleep(time.Second * 3)
	fmt.Println(a)
}
