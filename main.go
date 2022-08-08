package main

import (
	"flag"
	"fmt"
	"github.com/ungame/go-stats/stats"
	"strconv"
	"strings"
)

var numbers string

func init() {
	flag.StringVar(&numbers, "n", "", "set numbers separated with comma, ex: 1,2,3,4,5")
	flag.Parse()
}

func main() {
	values, err := parse(numbers)
	if err != nil {
		panic(err)
	}
	rol, err := stats.Rol(stats.ORDER_BY_ASC, values...)
	if err != nil {
		panic(err)
	}
	fmt.Println("Rol:            ", rol)
	fmt.Println("Mean:           ", stats.Mean(rol))
	fmt.Println("Median:         ", stats.Median(rol))
	fmt.Println("Mode:           ", stats.Mode(rol))
	fmt.Println("Deviation:      ", stats.Deviation(rol))
	fmt.Println("Variance:       ", stats.Variance(rol))
	fmt.Println("Std. Deviation: ", stats.StandardDeviation(rol))
}

func parse(n string) ([]float64, error) {
	items := strings.Split(strings.TrimSpace(n), ",")
	values := make([]float64, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		value, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return values, err
		}
		values = append(values, value)
	}
	if len(values) < 2 {
		return nil, fmt.Errorf("input should have two or more numbers separated with comma, but received: %v", n)
	}
	return values, nil
}
