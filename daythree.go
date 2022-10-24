package main

import (
	"fmt"
)

func daythree() {
	fmt.Println("Day Three GO!")

	var data string = ReadFile("./data/day3/data.rawr")
	//fmt.Println(data)

	//var lines []string = strings.Split(data, "\n")
	//fmt.Println(lines)

	var X, Y = 0, 0
	var houses = make(map[string]int)

	//star 2
	var isSanta bool = true
	var XS, YS, XR, YR = 0, 0, 0, 0
	var housesRobot = make(map[string]int)

	//deleiver to first house
	var position = fmt.Sprintf("%v_%b", X, Y)

	houses[position] = 1
	fmt.Println(X, Y, houses)

	//star 2
	housesRobot[position] = 2
	fmt.Println(XS, YS, XR, YR, housesRobot)

	for _, direction := range data {
		//fmt.Println(direction)
		//fmt.Printf("%c\n", direction)
		switch direction {
		case '^':
			{
				//fmt.Println("gone up")
				X = X + 1
				if isSanta {
					XS = XS + 1
				} else {
					XR = XR + 1
				}

			}
		case 'v':
			{
				//fmt.Println("gone down")
				X = X - 1

				if isSanta {
					XS = XS - 1
				} else {
					XR = XR - 1
				}
			}
		case '<':
			{
				//fmt.Println("gone left")
				Y = Y - 1
				if isSanta {
					YS = YS - 1
				} else {
					YR = YR - 1
				}
			}
		case '>':
			{
				//fmt.Println("gone right")
				Y = Y + 1

				if isSanta {
					YS = YS + 1
				} else {
					YR = YR + 1
				}
			}
		}

		position = fmt.Sprintf("%v_%b", X, Y)
		houses[position] = houses[position] + 1

		if isSanta {
			var santaPosition = fmt.Sprintf("%v_%b", XS, YS)
			housesRobot[santaPosition] = housesRobot[santaPosition] + 1
		} else {
			var robotPosition = fmt.Sprintf("%v_%b", XR, YR)
			housesRobot[robotPosition] = housesRobot[robotPosition] + 1
		}
		isSanta = !isSanta
	}

	var totalHouseVisted = len(houses)
	fmt.Printf("Star one result : %v\n", totalHouseVisted)

	var totalHouseRobotVisted = len(housesRobot)
	fmt.Printf("Star two result : %v\n", totalHouseRobotVisted)
}
