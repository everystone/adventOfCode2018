package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"

	chart "github.com/wcharczuk/go-chart"
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
	Members map[string]*member `json:"members"`
}

func getTime(ts string) time.Time {
	i, _ := strconv.ParseInt(ts, 10, 64)
	return time.Unix(i, 0)
}

func draw(values []chart.Value) {
	// https://github.com/wcharczuk/go-chart/blob/master/_examples/bar_chart/main.go
	bar := chart.BarChart{
		Title:      "Average time spent on part 2, in minutes.",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 100,
		XAxis:    chart.StyleShow(),
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Bars: values,
	}

	buffer := bytes.NewBuffer([]byte{})
	bar.Render(chart.PNG, buffer)
	ioutil.WriteFile("./charts/chart.png", buffer.Bytes(), 0644)
}

func main() {

	b, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var m members
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	// add extra whitespaces to short names, for cleaner output.
	for _, m := range m.Members {
		for len(m.Name) < 20 {
			m.Name += " "
		}
	}

	fmt.Printf("Name\t\t\tDay\tTime\t\tP1\t\tP2\n")
	fmt.Printf("--------------------------------------------------------------------\n")
	times := make(map[string][]float64)
	var chartValues []chart.Value
	for _, mem := range m.Members {

		// go map iteration is random, so sort first.
		var keys []int
		for day := range mem.Completion {
			keys = append(keys, day)
		}
		sort.Ints(keys)

		for _, day := range keys {
			v := mem.Completion[day]
			if len(v) == 2 {
				ts1 := getTime(v[1].Timestamp)
				ts2 := getTime(v[2].Timestamp)
				result := ts2.Sub(ts1).Minutes()
				fmt.Printf("%v\t%v\t%.2f min\t(%v -> %v)\n", mem.Name, day, result, ts1.Format("15:04:05"), ts2.Format("15:04:05"))
				times[mem.Name] = append(times[mem.Name], result)
			}
		}

		if len(times[mem.Name]) > 0 {
			var sum float64
			for _, i := range times[mem.Name] {
				sum += i
			}
			avg := sum / float64(len(times[mem.Name]))
			fmt.Printf("%v\tavg: \t%.2f min\n", mem.Name, avg)
			fmt.Printf("--------------------------------------------------------------------\n")

			if avg > 100 {
				avg = 100
			}
			chartValues = append(chartValues, chart.Value{Value: avg, Label: mem.Name})
		}
		sort.Slice(chartValues, func(i, j int) bool {
			return chartValues[i].Value > chartValues[j].Value
		})
		draw(chartValues)
	}
	// fmt.Printf("%v", m.Members["372116"])
}
