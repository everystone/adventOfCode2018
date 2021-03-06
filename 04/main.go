package main

import (
	"adventOfCode2018/common"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type guard struct {
	id      string
	slept   int
	minutes map[int]int
}

func getMin(line string) (int, error) {
	s := strings.Split(line, " ")
	t := strings.TrimRight(s[1], "]")
	m := strings.Split(t, ":")
	return strconv.Atoi(m[1])
}

func parseGuards(lines []string) map[string]*guard {
	guards := make(map[string]*guard)
	gid := ""
	sleepIndex := 0
	for i, l := range lines {
		s := strings.Split(l, " ")
		status := s[2]
		if strings.Contains(status, "Guard") {
			gid = s[3]
			if _, ok := guards[gid]; !ok {
				guards[gid] = &guard{
					id:      gid,
					slept:   0,
					minutes: make(map[int]int),
				}
			}
		}
		if strings.Contains(status, "wakes") {
			from, _ := getMin(lines[sleepIndex])
			to, _ := getMin(lines[i])
			guards[gid].slept += to - from
			for j := from; j < to; j++ {
				guards[gid].minutes[j]++
			}
		}
		if strings.Contains(status, "falls") {
			sleepIndex = i
		}
	}
	return guards
}

func part1(guards map[string]*guard) int {
	max, minute := 0, 0
	var sleeper *guard
	for _, g := range guards {
		logrus.Debugf("%v slept %v min", g.id, g.slept)
		if g.slept > max {
			max = g.slept
			sleeper = g
		}
	}
	logrus.Infof("Most sleepy guard: %v (%v minutes)", sleeper.id, sleeper.slept)
	max, minute = 0, 0
	for k, v := range sleeper.minutes {
		if v > max {
			max = v
			minute = k
		}
	}
	id, _ := strconv.Atoi(strings.Split(sleeper.id, "#")[1])
	logrus.Infof("Minute most slept: %v (%v times)", minute, max)
	return id * minute // 87681
}

func part2(guards map[string]*guard) int {
	max, minute := 0, 0
	var gg guard
	for _, g := range guards {
		for m, v := range g.minutes {
			if v > max {
				max = v
				minute = m
				gg = *g
			}
		}
	}
	id, _ := strconv.Atoi(strings.Split(gg.id, "#")[1])
	return id * minute // 136461
}

func main() {
	lines := common.ReadLines("./input.txt")
	sort.Strings(lines)
	guards := parseGuards(lines)
	p1 := part1(guards)
	logrus.Infof("Part 1 : %v", p1)
	p2 := part2(guards)
	logrus.Infof("Part 2: %v", p2)
}
