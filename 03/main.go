package main

import (
	"adventOfCode2018/common"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type claim struct {
	id, x1, y1, x2, y2 int
}

func main() {
	defer common.TimeTrack(time.Now(), "main")
	lines := common.ReadLines("./input.txt")
	var claims []claim
	for _, line := range lines {
		var x1, x2, y1, y2, id int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x1, &y1, &x2, &y2)
		x2 += x1
		y2 += y1
		claim := claim{id, x1, y1, x2, y2}
		claims = append(claims, claim)
		// logrus.Info(claim)
	}
	gridWidth := 1000
	exists := struct{}{}
	overlaps := make(map[int]struct{})
	lookup := make(map[int]int)
	grid := make([]int, gridWidth*gridWidth)
	for _, t := range claims {
		for x := t.x1; x < t.x2; x++ {
			for y := t.y1; y < t.y2; y++ {
				index := y*gridWidth + x
				grid[index]++
				if grid[index] >= 2 {
					overlaps[t.id] = exists
					overlaps[lookup[index]] = exists // previous claim is also overlapping.
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
	logrus.Infof("Part 1: %v", sum)
	for _, t := range claims {
		if _, ok := overlaps[t.id]; !ok {
			logrus.Infof("Part 2: #%v", t.id)
			return
		}
	}
}
