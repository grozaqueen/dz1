package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scan(&a)
	var b float64
	fmt.Scan(&b)
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
}
