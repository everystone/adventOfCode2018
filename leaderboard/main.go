package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Star struct {
	Timestamp string `json:"get_star_ts"`
}

type Member struct {
	LocalScore int `json:"local_score"`
	//LastStarTs  string `json:"last_star_ts"`
	Stars       int                  `json:"stars"`
	GlobalScore int                  `json:"global_score"`
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Completion  map[int]map[int]Star `json:"completion_day_level"`
}

type Members struct {
	Members map[string]Member `json:"members"`
}

func main() {

	byteValue, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//fmt.Printf("%v", byteValue)
	var m Members
	err = json.Unmarshal(byteValue, &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 2 completion times:\n")
	times := make(map[string][]float64)
	for _, mem := range m.Members {
		hasData := false

		for day, v := range mem.Completion {
			if len(v) == 2 {
				hasData = true
				i, _ := strconv.ParseInt(v[1].Timestamp, 10, 64)
				ts1 := time.Unix(i, 0)
				i, _ = strconv.ParseInt(v[2].Timestamp, 10, 64)
				ts2 := time.Unix(i, 0)
				result := ts2.Sub(ts1).Minutes()
				fmt.Printf("%v day %v p2: %.2f minutes\n", mem.Name, day, result)
				times[mem.Name] = append(times[mem.Name], result)
			}
		}

		if hasData {
			var sum float64
			for _, i := range times[mem.Name] {
				sum += i
			}
			fmt.Printf("%v avg part2 time: %.2f minutes\n", mem.Name, sum/float64(len(times[mem.Name])))
		}
	}
	// fmt.Printf("%v", m.Members["372116"])
}
