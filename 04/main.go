package main

import (
	"adventOfCode2018/common"
	"fmt"
	"log"
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
	guards := make(map[string]*guard)
	sort.Strings(lines)
	gid := ""
	wakeIndex := 0
	sleepIndex := 0
	for i, l := range lines {
		//fmt.Printf("%v %v\n", i, l)
		s := strings.Split(l, " ")
		// minute := strings.TrimRight(s[1], "]")
		status := s[2]
		if strings.Contains(status, "Guard") {
			gid = s[3]
			wakeIndex = -1
			sleepIndex = -1
			if _, ok := guards[gid]; !ok {
				guard := guard{
					id:    gid,
					slept: 0,
				}
				guards[gid] = &guard
			}
			fmt.Printf("new guard: %v\n", gid)
		}
		if strings.Contains(status, "wakes") {
			wakeIndex = i
			from, _ := getMin(lines[sleepIndex])
			to, _ := getMin(lines[wakeIndex])
			guards[gid].slept += to - from
			for j := from; j < to; j++ {
				guards[gid].minutes = append(guards[gid].minutes, j)
			}
			fmt.Printf("guard %v wakes up (%v - %v)\n", gid, to, from)

		}
		if strings.Contains(status, "falls") {
			sleepIndex = i
		}

	}

	max := 0
	var sleeper *guard
	for k, v := range guards {
		log.Printf("%v slept %v min", k, v.slept)
		if v.slept > max {
			max = v.slept
			sleeper = v
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
}
