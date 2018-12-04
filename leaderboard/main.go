package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	for _, v := range m.Members {
		log.Printf("%v", v.Name)
		for day, v := range v.Completion {
			log.Printf("Day %v", day)
			for _, s := range v {
				i, _ := strconv.ParseInt(s.Timestamp, 10, 64)
				ts := time.Unix(i, 0)
				log.Printf("%v", ts)
			}
		}
	}
	// fmt.Printf("%v", m.Members["372116"])

}
