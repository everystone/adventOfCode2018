package main

import (
	"adventOfCode2018/common"
	"fmt"
)

func main() {

	lines := common.ReadLines("./input.txt")
	//s := stack.New()
	m := make(map[string]string)
	for _, l := range lines {
		a := string(l[5])
		b := string(l[36])
		m[a] = b
		fmt.Printf("%v before %v\n", a, b)
	}
}
