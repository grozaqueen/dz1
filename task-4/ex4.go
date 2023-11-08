package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	var flag = 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if a[i][j] != a[j][i] {
				flag = 1
			}
		}
	}
	if flag == 0 {
		fmt.Print("yes")
	} else {
		fmt.Print("no")
	}
}
