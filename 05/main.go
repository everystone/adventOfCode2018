package main

import (
	"adventOfCode2018/common"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// dabAcCaCBAcCcaDA
func react(str string) (bool, string) {
	l := len(str)
	for i, c := range str {
		if i == l-1 {
			return false, str
		}
		s := string(c)
		if s == strings.ToLower(s) && string(str[i+1]) == strings.ToUpper(s) ||
			s == strings.ToUpper(s) && string(str[i+1]) == strings.ToLower(s) {
			// remove current + next
			match := str[i : i+2]
			//logrus.Infof("match: %v", match)
			return true, strings.Replace(str, match, "", -1)
		}
	}
	return false, str
}

func process(str string, unit string, results map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	running := true
	if unit != "" {
		str = strings.Replace(str, strings.ToLower(unit), "", -1)
		str = strings.Replace(str, strings.ToUpper(unit), "", -1)
	}
	for running == true {
		running, str = react(str)
	}
	logrus.Infof("result of unit %v: %v", unit, len(str))
	results[unit] = len(str)
}

func main() {
	defer common.TimeTrack(time.Now(), "main")
	input := common.ReadLines("./input.txt")[0]
	results := make(map[string]int)
	letters := []string{"", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "s", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var wg sync.WaitGroup
	for _, l := range letters {
		wg.Add(1)
		go process(input, l, results, &wg)
	}
	wg.Wait()
	min := len(input)
	best := ""
	for k, v := range results {
		if v < min {
			min = v
			best = k
		}
	}
	logrus.Infof("Part 1: %v", results[""])    // 9296
	logrus.Infof("Part 2: %v (%v)", min, best) // 5534, o
	// 28 seconds -> 8 seconds after implementing goroutines & syncgroup.
}
