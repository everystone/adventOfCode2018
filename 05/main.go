package main

import (
	"adventOfCode2018/common"
	"strings"

	"github.com/sirupsen/logrus"
)

// dabAcCaCBAcCcaDA
func process(str string, remove string) (bool, string) {
	if remove != "" {
		str = strings.Replace(str, strings.ToLower(remove), "", -1)
		str = strings.Replace(str, strings.ToUpper(remove), "", -1)
	}
	l := len(str)
	for i, c := range str {
		if i == l-1 {
			return false, str
		}
		s := string(c)

		// logrus.Infof("checking %v & %v", strings.ToLower(s), strings.ToUpper(s))
		if s == strings.ToLower(s) && string(str[i+1]) == strings.ToUpper(s) ||
			s == strings.ToUpper(s) && string(str[i+1]) == strings.ToLower(s) {
			// remove current + next
			match := str[i : i+2]
			//logrus.Infof("match: %v", match)
			return true, strings.Replace(str, match, "", -1)

		}
	}
	// for every char, check if the next one inverts it (if lower, next is upper, or opposite)
	return false, str
}

func main() {

	raw := common.ReadLines("./input.txt")[0]
	input := raw

	running := true
	for running == true {
		running, input = process(input, "")
	}

	logrus.Infof("Part 1: %v", len(input)) // 9296

	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "s", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	min := len(raw)
	best := ""
	for _, l := range letters {

		input = raw
		running = true
		for running == true {
			running, input = process(input, l)
		}
		result := len(input)
		if result < min {
			min = result
			best = l
		}
		logrus.Infof("result of unit %v: %v", l, result)

	}
	logrus.Infof("Part 2: %v (%v)", min, best) // 5534, o
}
