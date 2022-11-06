package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayten() {

	fmt.Println("Day Ten GO!")

	var data string = ReadFile("./data/day10/data.rawr")

	//var lines []string = strings.Split(data, "\n")
	fmt.Printf("Data : %v\n", data)
	var Output string = data

	for i := 0; i < 40; i++ {
		Output = lookSay(Output)
	}

	var starOne = len(Output)

	Output = data

	for i := 0; i < 50; i++ {
		Output = lookSay(Output)
	}

	var starTwo = len(Output)

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func lookSay(startingPoint string) string {
	var splitData = strings.Split(startingPoint, "")

	var currentValue = "0"
	var count = 0
	var newValue []string
	for i := range splitData {
		//fmt.Printf("Data char : %v\n", splitData[i])
		if currentValue == "0" {
			currentValue = splitData[i]
			count++
		} else if currentValue == splitData[i] {
			count++

		} else if i >= len(splitData) {
			newValue = append(newValue, strconv.Itoa(count))
			newValue = append(newValue, splitData[i])
			currentValue = splitData[i]
			count = 1
		} else {
			newValue = append(newValue, strconv.Itoa(count))
			newValue = append(newValue, currentValue)
			currentValue = splitData[i]
			count = 1
		}

	}

	if count != 0 {
		newValue = append(newValue, strconv.Itoa(count))
		newValue = append(newValue, currentValue)
	}

	return strings.Join(newValue, "")
}
