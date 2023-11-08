package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a+c > b && a+b > c && c+b > a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
