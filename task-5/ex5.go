package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	s = strings.Replace(s, "1", "one", len(s))
	fmt.Print(s)
}
