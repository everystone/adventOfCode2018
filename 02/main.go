package main

import (
	"adventOfCode2018/common"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	lines := common.ReadLines("./input.txt")
	total := make(map[int]int)
	for _, line := range lines {
		chars := common.Str2map(line)
		two := false
		three := false

		for _, val := range chars {
			if val == 2 && !two {
				total[2]++
				two = true
			}
			if val == 3 && !three {
				total[3]++
				three = true
			}
		}
	}
	log.Infof("part 1: %v\n", total[2]*total[3])

	for _, id := range lines {
		for _, id2 := range lines {
			if id == id2 {
				continue
			}
			misses := 0
			for i, v := range []byte(id) {
				if id2[i] != v {
					misses++
				}
			}
			if misses == 1 {
				var ans []string
				for i, v := range id {
					if id2[i] == byte(v) {
						ans = append(ans, string(v))
					}
				}
				log.Infof("part 2: %v\n", strings.Join(ans, ""))
				return
			}
		}
	}
}
