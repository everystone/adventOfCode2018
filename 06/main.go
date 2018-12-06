package main

import (
	"adventOfCode2018/common"
	"fmt"
	"log"
	"math"
)

type coordinate struct {
	x       int
	y       int
	size    int
	trapped bool
}

func (p *coordinate) String() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func manhattan(p1 coordinate, p2 coordinate) int {
	res := math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y))
	return int(res)
}

var coordinates []*coordinate
var grid [400][400]*coordinate

func main() {

	lines := common.ReadLines("./input.txt")
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		p := coordinate{x, y, 0, false}
		grid[p.x][p.y] = &p
		coordinates = append(coordinates, &p)
	}
	fmt.Printf("num coordinates: %v\n", len(coordinates))
	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			pos := coordinate{x, y, 0, false}

			min := 400
			var c *coordinate
			for _, coordinate := range coordinates {
				dist := manhattan(*coordinate, pos)
				if dist < min {
					c = coordinate
					min = dist
				}
			}
			grid[x][y] = c
			c.size++
			// fmt.Printf("%v belongs to %v\n", pos, c)
		}
	}
	// calc trapped
	for _, c1 := range coordinates {
		left, right, top, bot := false, false, false, false
		// check horizontal
		for x := 0; x < 400; x++ {
			c2 := grid[x][c1.y]
			if c1 == c2 {
				continue
			}
			if x < c1.x {
				left = true
			}
			if x > c1.x {
				right = true
			}
		}
		// check vertical
		for y := 0; y < 400; y++ {
			c2 := grid[c1.x][y]
			if c1 == c2 {
				continue
			}
			if y > c1.y {
				top = true
			}
			if y < c1.y {
				bot = true
			}
		}
		c1.trapped = left && right && bot && top
	}

	max, numtrapped := 0, 0
	var largest *coordinate
	for _, c := range coordinates {
		if c.trapped {
			numtrapped++
			if c.size > max {
				max = c.size
				largest = c
			}
		}
	}
	log.Printf("num trapped : %v\n", numtrapped)
	log.Printf("largest trapped: %v: %v", largest, max) // 5626

	/*
		What is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000?
	*/

	var safe []*coordinate

	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			c1 := &coordinate{x, y, 0, false}
			total := 0
			for _, c2 := range coordinates {
				if c1 == c2 {
					continue
				}
				dist := manhattan(*c1, *c2)
				total += dist
			}
			if total < 10000 {
				safe = append(safe, c1)
			}
		}
	}

	fmt.Printf("safe: %v\n", len(safe)) // 46554
}
