package main

import (
	"advent2018/common"
	"fmt"
)

func main() {

	lines := common.ReadLines("./input.txt")
	total := make(map[int]int)
	for _, line := range lines {
		chars := common.Str2map(line)
		two := false
		three := false

		for _, val := range chars {
			if val == 2 && !two {
				total[2]++
				two = true
			}
			if val == 3 && !three {
				total[3]++
				three = true
			}
		}
	}
	fmt.Printf("%v\n", total)
	fmt.Printf("sum: %v\n", total[2]*total[3])

	for _, id := range lines {
		for _, id2 := range lines {
			if id == id2 {
				continue
			}
			misses := 0
			var m []string
			for i, v := range []byte(id) {
				if id2[i] != v {
					misses++
					m = append(m, string(v))
				}
			}
			if misses < 2 {
				fmt.Printf("%s -> \n%s: %v: %v\n\n", id, id2, misses, m)
			}
		}
	}

}
