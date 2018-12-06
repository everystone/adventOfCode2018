package main

import (
	"adventOfCode2018/common"
	"fmt"
	"math"
)

type coord struct {
	x    int
	y    int
	inf  bool
	size int
}

func (p *coord) String() string {
	return "A"
}

func getDistance(p1 *coord, p2 *coord) int {
	res := math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y))
	return int(res)
}

var coordinates []*coord
var grid [300][300]*coord

func main() {

	lines := common.ReadLines("./input.txt")
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		p := coord{x, y, false, 0}
		grid[p.x][p.y] = &p
		coordinates = append(coordinates, &p)
	}
	fmt.Printf("num coordinates: %v\n", len(coordinates))

	// loop grid
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			// find coord closest to i,j
			dist := getDistance(p1, p2)
			fmt.Printf("distance: %v\n", dist)
		}
	}
	num := 0
	for _, p := range coordinates {
		if p.inf {
			num++
		}
	}
	fmt.Printf("inf count: %v\n", num)
}
