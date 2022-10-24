package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func daytwo() {
	fmt.Println("Day Two GO!")

	var data string = ReadFile("./data/day2/data.rawr")
	//fmt.Println(data)

	var lines []string = strings.Split(data, "\n")

	var presents []int64
	var ribbons []int64

	for _, line := range lines {
		var values []string = strings.Split(line, `x`)
		var parsed [3]int64
		for index, value := range values {
			i, _ := strconv.ParseInt(value, 0, 64)
			parsed[index] = i
		}

		sides := [3]int64{2 * parsed[0] * parsed[1], 2 * parsed[0] * parsed[2], 2 * parsed[1] * parsed[2]}

		sort.Slice(sides[:], func(i, j int) bool {
			return sides[i] < sides[j]
		})

		var dimensions int64 = sides[0] + sides[1] + sides[2]

		var littleBitExtra = sides[0] / 2

		presents = append(presents, (dimensions + littleBitExtra))

		//star two
		sort.Slice(parsed[:], func(i, j int) bool {
			return parsed[i] < parsed[j]
		})

		var ribbonlength = ((parsed[0] * 2) + (parsed[1] * 2))
		var bow = parsed[0] * parsed[1] * parsed[2]
		ribbons = append(ribbons, ribbonlength+bow)
	}

	var totalWrapingPaper = sum64(presents)
	//fmt.Println(totalWrapingPaper)
	fmt.Printf("Star one result : %v\n", totalWrapingPaper)

	var totalRibbon int64 = sum64(ribbons)
	fmt.Printf("Star two result : %v\n", totalRibbon)
}
