package common

import (
	"sort"
	"strings"
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
