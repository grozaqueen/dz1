package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	m := a[0]
	for i := 1; i < n; i++ {
		a[i-1] = a[i]
	}
	a[n-1] = m
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
}
