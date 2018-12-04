package main

import (
	"adventOfCode2018/common"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type claim struct {
	id int
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	// defer common.TimeTrack(time.Now(), "main")
	lines := common.ReadLines("./input.txt")
	var claims []claim
	for _, line := range lines {
		s := strings.Split(line, " ")
		idStr, posStr, sizeStr := s[0], s[2], s[3]
		posStr = strings.TrimRight(posStr, ":")
		size := strings.Split(sizeStr, "x")
		pos := strings.Split(posStr, ",")
		x1, _ := strconv.Atoi(pos[0])
		y1, _ := strconv.Atoi(pos[1])
		x2, _ := strconv.Atoi(size[0])
		y2, _ := strconv.Atoi(size[1])
		x2 += x1
		y2 += y1
		id, _ := strconv.Atoi(strings.Split(idStr, "#")[1])
		claims = append(claims, claim{id, x1, y1, x2, y2})
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
