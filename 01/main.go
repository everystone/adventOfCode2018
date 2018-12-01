package main

import (
	"advent2018/common"
	"log"
	"strconv"
)

func main() {
	fileName := "./input.txt"
	lines := common.ReadLines(fileName)
	freq := make(map[int]struct{})
	var exists = struct{}{}

	sum := 0
	for {
		for _, s := range lines {
			op := string(s[0])
			val, _ := strconv.Atoi(s[1:])
			// log.Printf("%v = %v %v", s, op, val)
			switch op {
			case "+":
				sum += val
			case "-":
				sum -= val
			}

			// check if value has been seen before
			if _, ok := freq[sum]; ok {
				log.Printf("found duplicate freq: %v", sum)
				return
			}
			freq[sum] = exists
		}
		log.Printf("sum: %v", sum)
	}
}
