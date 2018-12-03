package main

import (
	"adventOfCode2018/common"
	"log"
	"strconv"
	"strings"
)

type claim struct {
	id string
	x  int
	y  int
	w  int
	h  int
}

func main() {
	lines := common.ReadLines("./input.txt")
	var claims []claim
	for _, line := range lines {
		s := strings.Split(line, " ")
		id, posStr, sizeStr := s[0], s[2], s[3]
		posStr = strings.TrimRight(posStr, ":")
		pos := strings.Split(posStr, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		size := strings.Split(sizeStr, "x")
		w, _ := strconv.Atoi(size[0])
		h, _ := strconv.Atoi(size[1])
		claims = append(claims, claim{id, x, y, w, h})
	}
	gridWidth := 1000
	overlaps := make(map[string]bool)
	lookup := make(map[int]string)
	grid := make([]int, gridWidth*gridWidth)
	for _, t := range claims {
		lx1, ly1, lx2, ly2 := t.x, t.y, t.x+t.w, t.y+t.h
		// create grid to count overlaps of each square.
		for x := lx1; x < lx2; x++ {
			for y := ly1; y < ly2; y++ {
				index := y*gridWidth + x
				grid[index]++
				if grid[index] >= 2 {
					overlaps[t.id] = true
					overlaps[lookup[index]] = true // previous claim is also overlapping.
				}
				lookup[index] = t.id
			}
		}
	}
	sum := 0 // 113966
	for _, val := range grid {
		if val >= 2 {
			sum++
		}
	}
	log.Printf("Part 1: %v", sum)
	for _, t := range claims {
		if _, ok := overlaps[t.id]; !ok {
			log.Printf("Part 2: %v", t.id)
		}
	}
}
