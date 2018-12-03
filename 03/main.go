package main

import (
	"advent2018/common"
	"log"
	"strconv"
	"strings"
)

func collides(x int, y int) bool {
	return true
}

type tile struct {
	id string
	x  int
	y  int
	w  int
	h  int
}

func main() {
	lines := common.ReadLines("./input.txt")
	var tiles []tile
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

		tiles = append(tiles, tile{id, x, y, w, h})
		//log.Printf("id %s at %v,%v size %vx%v", id, x, y, w, h)
	}
	log.Printf("found %v tiles.", len(tiles))
	collisions := 0
	for _, t := range tiles {
		for _, t2 := range tiles {
			if t.id == t2.id {
				continue
			}
			// check if t2 overlaps t1
			lx1, ly1, lx2, ly2 := t.x, t.y, t.x+t.w, t.y-t.h
			rx1, ry1, rx2, ry2 := t2.x, t2.y, t2.x+t2.w, t2.y-t2.h

			if lx1 < rx2 && lx2 > rx1 && ly1 > ry2 && ly2 < ry1 {
				collisions++
				log.Printf("overlap: %v %v", t, t2)
				log.Printf("%v %v", rx2-lx1, ry2-ly1)
				return
			}
		}
	}
	log.Printf("collisions: %v", collisions)
}
