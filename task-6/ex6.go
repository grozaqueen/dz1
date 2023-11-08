package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var myMap map[string]int
	myMap = make(map[string]int)
	var kol int = 0
	var s string
	if reader.Scan() {
		s = reader.Text()
	}
	a := strings.Split(s, " ")
	for i := 0; i < len(a); i++ {
		myMap[a[i]]++
		if myMap[a[i]] == 1 {
			kol++
		}
	}
	fmt.Print(kol)
}
