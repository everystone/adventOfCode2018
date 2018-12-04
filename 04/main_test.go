package main

import (
	"adventOfCode2018/common"
	"sort"
	"testing"
	"time"
)

func TestParseGuards(t *testing.T) {
	lines := common.ReadLines("./input.txt")
	sort.Strings(lines)
	guards := parseGuards(lines)
	g := guards["#3559"]
	if g.id != "#3559" {
		t.Error("Failed to parse id")
	}
	if len(g.minutes) != 59 {
		t.Errorf("Failed to parse minutes: %v", len(g.minutes))
	}
	if g.slept != 410 {
		t.Errorf("Failed to parse slept: %v", g.slept)
	}
}

func TestPart1(t *testing.T) {
	lines := common.ReadLines("./input.txt")
	sort.Strings(lines)
	guards := parseGuards(lines)
	p1 := part1(guards)
	if p1 != 87681 {
		t.Errorf("Part1 failed: actual %v expected %v", p1, 87681)
	}
}

func TestPart2(t *testing.T) {
	lines := common.ReadLines("./input.txt")
	sort.Strings(lines)
	guards := parseGuards(lines)
	p2 := part2(guards)
	if p2 != 136461 {
		t.Errorf("Part1 failed: actual %v expected %v", p2, 136461)
	}
}

func TestPerformance(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		main()
	}
	elapsed := time.Since(start)
	t.Logf("average speed %s", elapsed/100)
}
