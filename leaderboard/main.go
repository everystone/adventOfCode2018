package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type star struct {
	Timestamp string `json:"get_star_ts"`
}

type member struct {
	LocalScore int `json:"local_score"`
	//LastStarTs  string `json:"last_star_ts"`
	Stars       int                  `json:"stars"`
	GlobalScore int                  `json:"global_score"`
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Completion  map[int]map[int]star `json:"completion_day_level"`
}

type members struct {
	Members map[string]member `json:"members"`
}

func main() {

	byteValue, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//fmt.Printf("%v", byteValue)
	var m members
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
				fmt.Printf("%v day %v\t%.2f minutes\t(%v)\n", mem.Name, day, result, ts2.Format("15:04:05"))
				times[mem.Name] = append(times[mem.Name], result)
			}
		}

		if hasData {
			var sum float64
			for _, i := range times[mem.Name] {
				sum += i
			}
			fmt.Printf("%v avg:  \t%.2f minutes\n", mem.Name, sum/float64(len(times[mem.Name])))
			fmt.Printf("-----------------------------------------------\n")
		}
	}
	// fmt.Printf("%v", m.Members["372116"])
}
