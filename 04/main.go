package main

import (
	"adventOfCode2018/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type guard struct {
	id      string
	slept   int
	minutes []int
}

func getMin(line string) (int, error) {
	s := strings.Split(line, " ")
	t := strings.TrimRight(s[1], "]")
	m := strings.Split(t, ":")
	return strconv.Atoi(m[1])
}

func main() {
	lines := common.ReadLines("./input.txt")
	sort.Strings(lines)
	guards := make(map[string]*guard)
	gid := ""
	wakeIndex := 0
	sleepIndex := 0
	for i, l := range lines {
		s := strings.Split(l, " ")
		status := s[2]
		if strings.Contains(status, "Guard") {
			gid = s[3]
			wakeIndex = -1
			sleepIndex = -1
			if _, ok := guards[gid]; !ok {
				guards[gid] = &guard{
					id:    gid,
					slept: 0,
				}
			}
		}
		if strings.Contains(status, "wakes") {
			wakeIndex = i
			from, _ := getMin(lines[sleepIndex])
			to, _ := getMin(lines[wakeIndex])
			guards[gid].slept += to - from
			for j := from; j < to; j++ {
				guards[gid].minutes = append(guards[gid].minutes, j)
			}
			// fmt.Printf("guard %v wakes up (%v - %v)\n", gid, to, from)

		}
		if strings.Contains(status, "falls") {
			sleepIndex = i
		}

	}

	max := 0
	var sleeper *guard
	for _, g := range guards {
		// log.Printf("%v slept %v min", k, v.slept)
		if g.slept > max {
			max = g.slept
			sleeper = g
		}
	}
	fmt.Printf("sleeper: %v: %v", sleeper.id, sleeper.slept)
	ms := make(map[int]int)
	for _, v := range sleeper.minutes {
		ms[v]++
	}

	max = 0
	minute := 0
	for k, v := range ms {
		if v > max {
			max = v
			minute = k
		}
	}
	fmt.Printf("minute most slept: %v: %v\n", minute, max) // 87681

	// part 2
	type gs struct {
		guard   *guard
		minutes map[int]int
	}
	ms2 := make(map[*guard]*gs)
	max = 0
	minute = 0
	var gg guard
	for _, g := range guards {
		for _, v := range g.minutes {
			if _, ok := ms2[g]; !ok {
				ms2[g] = &gs{
					minutes: make(map[int]int),
				}
			}
			ms2[g].minutes[v]++
			if ms2[g].minutes[v] > max {
				max = ms2[g].minutes[v]
				minute = v
				gg = *g
			}
		}
	}
	id, _ := strconv.Atoi(strings.Split(gg.id, "#")[1])
	fmt.Printf("part 2: %v\n", id*minute)

}
