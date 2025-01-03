package main

import "fmt"

func absum1() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			return
		}
		fmt.Println(a + b)
	}
}
func absum2() {
	var n int

	var a, b int
	for {
		_, err1 := fmt.Scan(&n)
		if err1 != nil {
			return
		}
		for i := 0; i < n; i++ {
			_, err := fmt.Scan(&a, &b)
			if err != nil {
				break
			}
			fmt.Println(a + b)
		}
	}

}
func absum3() {
	var a, b int
	for {
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			return
		}
		fmt.Println(a + b)
	}
}
func absum4() {
	var n, a, sum int
	for {
		fmt.Scan(&n)
		if n == 0 {
			return
		}
		sum = 0
		for i := 0; i < n; i++ {
			fmt.Scan(&a)
			sum += a
		}
		fmt.Println(sum)
	}
}
func testSlice(a []int) {
	a[0] = 1
}
func main() {
	//absum4()
	a := make([]int, 5)
	testSlice(a)
	fmt.Println(a)
}
