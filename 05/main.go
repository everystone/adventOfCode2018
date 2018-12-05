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
		unit := strmap[s]
		if s == unit.lower && string(str[i+1]) == unit.upper ||
			s == unit.upper && string(str[i+1]) == unit.lower {
			// remove current + next
			match := str[i : i+2]
			// logrus.Infof("match: %v", match)
			return true, strings.Replace(str, match, "", -1)
		}
	}
	return false, str
}

func process(str string, unit string, results map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	running := true
	if unit != "" {
		u := strmap[unit]
		str = strings.Replace(str, u.lower, "", -1)
		str = strings.Replace(str, u.upper, "", -1)
	}
	for running == true {
		running, str = react(str)
	}
	logrus.Infof("result of unit %v: %v", unit, len(str))
	results[unit] = len(str)
}

type unit struct {
	upper string
	lower string
}

// cached string upper/lower case lookup, to reduce number of toUpper / toLower calls.
var strmap map[string]unit

func createUnitMap(letters []string) {
	for _, l := range letters {
		unit := unit{
			lower: l,
			upper: strings.ToUpper(l),
		}
		strmap[l] = unit
		strmap[unit.upper] = unit
	}
}

func main() {
	defer common.TimeTrack(time.Now(), "main")
	letters := []string{"", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	input := common.ReadLines("./input.txt")[0]
	strmap = make(map[string]unit)
	createUnitMap(letters)
	// fmt.Println(strmap)

	results := make(map[string]int)

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
	//  initial 28 seconds
	// -> 8 seconds after implementing goroutines & syncgroup.
	// -> 6 seconds after removing unit from string before react loop.
	// -> 3,8 seconds after caching upper/lower variants of chars in strmap.
}

/**
Showing top 10 nodes out of 38
      flat  flat%   sum%        cum   cum%
     6.56s 19.68% 19.68%     11.50s 34.50%  runtime.mallocgc
     3.27s  9.81% 29.49%     32.03s 96.10%  adventOfCode2018/05.react
     3.09s  9.27% 38.76%     11.67s 35.01%  strings.ToLower
		 2.64s  7.92% 46.68%     11.80s 35.40%  strings.ToUpper

after strmap:

Showing top 10 nodes out of 46
      flat  flat%   sum%        cum   cum%
    4960ms 25.53% 25.53%     8920ms 45.91%  runtime.mapaccess1_faststr
    3350ms 17.24% 42.77%    18630ms 95.88%  adventOfCode2018/05.react
    3290ms 16.93% 59.70%     3290ms 16.93%  runtime.memeqbody
    1950ms 10.04% 69.74%     2760ms 14.20%  runtime.intstring
    1210ms  6.23% 75.97%     1210ms  6.23%  runtime.aeshashbody
**/
