package common

import (
	"log"
	"sort"
	"strings"
	"time"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Str2map(s string) map[byte]int {
	chars := make(map[byte]int)
	for _, r := range []byte(s) {
		chars[r]++
	}
	return chars
}

// returns key from map where value is highest
func maxInt(m map[int]int) int {
	max, key := 0, 0
	for k, v := range m {
		if v > max {
			max = v
			key = k
		}
	}
	return key
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
