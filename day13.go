package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//	type requirement struct {
//		name  string
//		value int
//	}
type dinnerPerson struct {
	happiness    int
	requirements map[string]int
}

func daythirdteen() {

	fmt.Println("Day Thirdteen GO!")

	var data string = ReadFile("./data/day13/data.rawr")

	//fmt.Printf("Data : %v\n", data)
	var lines []string = strings.Split(data, "\n")

	var table map[string]map[string]int = make(map[string]map[string]int)
	var people []string

	for i := 0; i < len(lines); i++ {

		//var guest dinnerPerson
		splitRow := strings.Split(lines[i], " ")
		//guest = dinnerPerson{splitRow[0], 0, nil}
		neigbour := strings.Replace(splitRow[10], ".", "", -1)
		value, _ := strconv.Atoi(splitRow[3])

		if splitRow[2] == "lose" {
			value = -(value)
		}

		fmt.Printf("Guest %v : %v -  %v - %v\n", splitRow[0], splitRow[2], neigbour, value)
		if !slices.Contains(people, splitRow[0]) {
			//distrint List of people
			people = append(people, splitRow[0])
		}

		if table[splitRow[0]] == nil {
			table[splitRow[0]] = make(map[string]int)
		}
		table[splitRow[0]][neigbour] = value

	}

	allPermutations := permutations(people)
	var happinessCharges []int

	for i := 0; i < len(allPermutations); i++ {
		charges := 0
		for j := 0; j < len(allPermutations[i]); j++ {

			var neighboursRight string
			var neighboursLeft string

			if j == 0 {
				neighboursRight = allPermutations[i][j+1]
				neighboursLeft = allPermutations[i][len(allPermutations[i])-1]
			} else if j == len(allPermutations[i])-1 {
				neighboursLeft = allPermutations[i][j-1]
				neighboursRight = allPermutations[i][0]
			} else {
				neighboursRight = allPermutations[i][j+1]
				neighboursLeft = allPermutations[i][j-1]
			}

			//fmt.Printf("neighbours : %v - %v \n", neighboursLeft, neighboursRight)
			charges = charges + table[allPermutations[i][j]][neighboursLeft] + table[allPermutations[i][j]][neighboursRight]

		}
		happinessCharges = append(happinessCharges, charges)
	}

	//fmt.Printf("People : %v \n happinessCharges : %v \n", people, happinessCharges)
	//fmt.Printf("table : %v \n", table)

	sort.Slice(happinessCharges[:], func(i, j int) bool {
		return happinessCharges[i] > happinessCharges[j]
	})

	var starOne = happinessCharges[0]

	// add myself

	table["mySelf"] = make(map[string]int)
	for guestId := range people {
		table["mySelf"][people[guestId]] = 0
		table[people[guestId]]["mySelf"] = 0
	}
	people = append(people, "mySelf")
	allStar2Permutations := permutations(people)

	fmt.Printf("people : %v \n", people)
	fmt.Printf("table : %v \n", table)

	var happinessChargesStar2 []int
	for i := 0; i < len(allStar2Permutations); i++ {
		charges := 0
		for j := 0; j < len(allStar2Permutations[i]); j++ {

			var neighboursRight string
			var neighboursLeft string

			if j == 0 {
				neighboursRight = allStar2Permutations[i][j+1]
				neighboursLeft = allStar2Permutations[i][len(allStar2Permutations[i])-1]
			} else if j == len(allStar2Permutations[i])-1 {
				neighboursLeft = allStar2Permutations[i][j-1]
				neighboursRight = allStar2Permutations[i][0]
			} else {
				neighboursRight = allStar2Permutations[i][j+1]
				neighboursLeft = allStar2Permutations[i][j-1]
			}

			//fmt.Printf("neighbours : %v - %v \n", neighboursLeft, neighboursRight)
			charges = charges + table[allStar2Permutations[i][j]][neighboursLeft] + table[allStar2Permutations[i][j]][neighboursRight]

		}
		happinessChargesStar2 = append(happinessChargesStar2, charges)
	}

	sort.Slice(happinessChargesStar2[:], func(i, j int) bool {
		return happinessChargesStar2[i] > happinessChargesStar2[j]
	})

	var starTwo = happinessChargesStar2[0]

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}
