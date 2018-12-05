package main

import (
	"adventOfCode2018/common"
	"strings"

	"github.com/sirupsen/logrus"
)

// dabAcCaCBAcCcaDA
func process(str string, m string, custom bool) (bool, string) {
	l := len(str)
	for i, c := range str {
		if i == l-1 {
			return false, str
		}
		s := string(c)
		if !custom {
			m = s
		}
		// logrus.Infof("checking %v & %v", strings.ToLower(m), strings.ToUpper(m))
		if s == strings.ToLower(m) && string(str[i+1]) == strings.ToUpper(m) ||
			s == strings.ToUpper(m) && string(str[i+1]) == strings.ToLower(m) {
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

	// logrus.Infof("len of str %v", len(input))
	running := true
	// for running == true {
	// 	running, input = process(input, "", false)
	// }
	// //logrus.Infof("%v", input)
	// logrus.Infof("len of str %v", len(input)) // 9296

	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	min := len(raw)
	best := ""
	for _, l := range letters {

		input = raw
		running = true
		// logrus.Infof("testing unit %v %v", l, len(input))
		for running == true {
			running, input = process(input, l, true)
		}
		result := len(input)
		if result < min {
			min = result
			best = l
		}
		logrus.Infof("result of unit %v: %v", l, result)

	}
	logrus.Infof("winner: %v (%v)", best, min)
}
