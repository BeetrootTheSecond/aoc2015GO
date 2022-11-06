package main

import (
	"fmt"
)

func dayone() {

	fmt.Println("Day One GO!")

	var data string = ReadFile("./data/day1/data.rawr")

	//fmt.Println(data)

	var floor int = 0

	var basementEntry *int

	fmt.Printf("Starting floor %v\n", floor)

	for index, char := range data {
		//fmt.Println(char)
		//fmt.Printf("%c\n", char)
		if char == 40 {
			floor = floor + 1
		} else if char == 41 {
			floor = floor - 1
		}

		if basementEntry == nil && floor == -1 {
			basementEntry = new(int)
			*basementEntry = index + 1
		}
	}

	fmt.Printf("Star one result : %v\n", floor)

	fmt.Printf("Star two result : %v\n", *basementEntry)
}
