package main

import (
	"adventOfCode2018/common"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/golang-collections/collections/stack"
	"github.com/sirupsen/logrus"
)

func react(str string) int {
	s := stack.New()
	for _, ch := range str {
		c := string(ch)
		if s.Len() == 0 {
			s.Push(c)
		} else {
			last := s.Peek()
			current := string(c)
			unit := strmap[current]
			if last == unit.lower && current == unit.upper ||
				last == unit.upper && current == unit.lower {
				s.Pop()
			} else {
				s.Push(c)
			}
		}
	}
	return s.Len()
}

func process(str string, unit string, results map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	if unit != "" {
		u := strmap[unit]
		str = strings.Replace(str, u.lower, "", -1)
		str = strings.Replace(str, u.upper, "", -1)
	}

	result := react(str)

	logrus.Debugf("result of unit %v: %v", unit, result)
	results[unit] = result
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
	fmt.Printf("Part 1: %v\n", results[""])    // 9296
	fmt.Printf("Part 2: %v (%v)\n", min, best) // 5534, o
	//  initial 28 seconds
	// -> 8 seconds after implementing goroutines & syncgroup.
	// -> 6 seconds after removing unit from string before react loop.
	// -> 3,3 seconds after caching upper/lower variants of chars in strmap.
	// -> ~100ms after implementing react() using a stack
}

/**
Duration: 6.17s, Total samples = 33.33s (540.46%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 27.18s, 81.55% of 33.33s total
Dropped 112 nodes (cum <= 0.17s)
Showing top 10 nodes out of 38
      flat  flat%   sum%        cum   cum%
     6.56s 19.68% 19.68%     11.50s 34.50%  runtime.mallocgc
     3.27s  9.81% 29.49%     32.03s 96.10%  adventOfCode2018/05.react
     3.09s  9.27% 38.76%     11.67s 35.01%  strings.ToLower
     2.64s  7.92% 46.68%     11.80s 35.40%  strings.ToUpper
     2.40s  7.20% 53.89%      2.40s  7.20%  runtime.acquirem (inline)
     2.38s  7.14% 61.03%      8.55s 25.65%  runtime.makeslice
		 2.30s  6.90% 67.93%      9.24s 27.72%  runtime.slicebytetostring

after strmap:

Duration: 3.74s, Total samples = 17.87s (477.68%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 16080ms, 89.98% of 17870ms total
Dropped 84 nodes (cum <= 89.35ms)
Showing top 10 nodes out of 48
      flat  flat%   sum%        cum   cum%
    4850ms 27.14% 27.14%     8440ms 47.23%  runtime.mapaccess1_faststr
    2890ms 16.17% 43.31%    16960ms 94.91%  adventOfCode2018/05.react
    2880ms 16.12% 59.43%     2880ms 16.12%  runtime.memeqbody
    1610ms  9.01% 68.44%     2440ms 13.65%  runtime.intstring
    1310ms  7.33% 75.77%     1310ms  7.33%  runtime.aeshashbody
     830ms  4.64% 80.41%      830ms  4.64%  runtime.encoderune
		 670ms  3.75% 84.16%      670ms  3.75%  runtime.memequal


after using stack_

Duration: 202ms, Total samples = 460ms (227.73%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 340ms, 73.91% of 460ms total
Showing top 10 nodes out of 57
      flat  flat%   sum%        cum   cum%
      70ms 15.22% 15.22%       70ms 15.22%  runtime.cgocall
      50ms 10.87% 26.09%       50ms 10.87%  runtime.(*gcWork).dispose
      50ms 10.87% 36.96%      140ms 30.43%  runtime.mallocgc
      40ms  8.70% 45.65%       40ms  8.70%  runtime.osyield
      30ms  6.52% 52.17%       30ms  6.52%  runtime.heapBitsForObject
      20ms  4.35% 56.52%      230ms 50.00%  adventOfCode2018/05.react
      20ms  4.35% 60.87%       20ms  4.35%  runtime.(*gcBits).bitp (inline)
**/
