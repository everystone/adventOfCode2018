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
	"github.com/wcharczuk/go-chart/drawing"
	"github.com/wcharczuk/go-chart/seq"
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

func drawBasic(times map[string][]float64, numDays int) {

	colors := []string{
		"FF99E6", "CCFF1A", "FF1A66", "E6331A", "33FFCC",
		"66994D", "B366CC", "4D8000", "B33300", "CC80CC",
		"66664D", "991AFF", "E666FF", "4DB3FF", "1AB399",
		"E666B3", "33991A", "CC9999", "B3B31A", "00E680",
		"4D8066", "809980", "E6FF80", "1AFF33", "999933",
		"FF3380", "CCCC00", "66E64D", "4D80CC", "9900B3",
		"E64D66", "4DB380", "FF4D4D", "99E6E6", "6666FF",
	}

	var series []chart.Series
	colorIndex := 0
	for k, v := range times {
		colorIndex++
		for i := len(v); i < numDays; i++ {
			v = append(v, 0)
		}

		color := drawing.ColorFromHex(colors[colorIndex])
		con := chart.ContinuousSeries{
			XValues: seq.Range(1.0, float64(numDays)),
			YValues: v,
			Name:    k,
			Style: chart.Style{
				Show:        true,
				StrokeColor: color,
				StrokeWidth: 2,
			},
		}
		series = append(series, con)
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Days",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Minutes",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: 60,
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 260,
			},
			//FillColor: drawing.ColorBlack,
		},
		Canvas: chart.Style{
			//FillColor: drawing.ColorBlack,
		},
		Series: series,
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	ioutil.WriteFile("./charts/basic.png", buffer.Bytes(), 0644)
}

func draw(title string, values []chart.Value, limit float64) {
	// https://github.com/wcharczuk/go-chart/blob/master/_examples/bar_chart/main.go

	bar := chart.BarChart{
		Title:      title,
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
			//FillColor: drawing.ColorBlack,
		},
		Height:   1024,
		BarWidth: 100,
		XAxis:    chart.StyleShow(),
		YAxis: chart.YAxis{
			Style:     chart.StyleShow(),
			Name:      "Names",
			NameStyle: chart.StyleShow(),
			// Range: &chart.ContinuousRange{
			// 	Min: 0,
			// 	Max: limit,
			// },
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
	var avgValues []chart.Value
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

		if len(times[mem.Name]) > 2 {
			var sum float64
			for _, i := range times[mem.Name] {
				sum += i
			}
			avg := sum / float64(len(times[mem.Name]))
			fmt.Printf("%v\tavg: \t%.2f min\n", mem.Name, avg)
			fmt.Printf("--------------------------------------------------------------------\n")

			avgValues = append(avgValues, chart.Value{Value: avg, Label: mem.Name})

		}
	}

	sort.Slice(avgValues, func(i, j int) bool {
		return avgValues[i].Value < avgValues[j].Value
	})
	draw("Top #5 part 2 average times", avgValues[0:5], 60)
	// draw("Total time (part 2)", sumValues, 500)
	//drawBasic(times, 12)
	// fmt.Printf("%v", m.Members["372116"])
}
