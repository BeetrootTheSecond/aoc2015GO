package main

import (
	"fmt"
	"strconv"
	"strings"
)

type aunt struct {
	number      int
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

var Sues []aunt

func daysixteen() {

	fmt.Println("Day Sixteen GO!")

	var data string = ReadFile("./data/day16/data.rawr")
	var lines []string = strings.Split(data, "\n")
	//fmt.Printf("%v \n", lines)

	for _, line := range lines {
		//fmt.Printf("%v \n", line)
		number := -1
		children := -1
		cats := -1
		samoyeds := -1
		pomeranians := -1
		akitas := -1
		vizslas := -1
		goldfish := -1
		trees := -1
		cars := -1
		perfumes := -1

		var properties = strings.Split(line, ", ")
		var sueSplit = strings.Split(properties[0], ":")
		number, _ = strconv.Atoi(strings.Replace(sueSplit[0], "Sue ", "", -1))
		sueSplit[1] = strings.Replace(sueSplit[1], " ", "", -1)
		properties[0] = strings.Join(sueSplit[1:], ":")

		for _, property := range properties {
			split := strings.Split(property, ": ")
			switch split[0] {
			case "children":
				{
					children, _ = strconv.Atoi(split[1])
					break
				}
			case "cats":
				{
					cats, _ = strconv.Atoi(split[1])
					break
				}
			case "samoyeds":
				{
					samoyeds, _ = strconv.Atoi(split[1])
					break
				}
			case "pomeranians":
				{
					pomeranians, _ = strconv.Atoi(split[1])
					break
				}
			case "akitas":
				{
					akitas, _ = strconv.Atoi(split[1])
					break
				}
			case "vizslas":
				{
					vizslas, _ = strconv.Atoi(split[1])
					break
				}
			case "goldfish":
				{
					goldfish, _ = strconv.Atoi(split[1])
					break
				}
			case "trees":
				{
					trees, _ = strconv.Atoi(split[1])
					break
				}
			case "cars":
				{
					cars, _ = strconv.Atoi(split[1])
					break
				}
			case "perfumes":
				{
					perfumes, _ = strconv.Atoi(split[1])
					break
				}

			}
		}

		ns := aunt{number, children, cats, samoyeds, pomeranians, akitas, vizslas, goldfish, trees, cars, perfumes}
		Sues = append(Sues, ns)
	}

	StarOneAunt := aunt{0, 3, 7, 2, 3, 0, 0, 5, 3, 2, 1}

	var FoundAunts []aunt
	var FoundAuntsStarTwo []aunt
	for _, Sue := range Sues {

		if Sue.children == StarOneAunt.children || Sue.children == -1 {
			if Sue.cats == StarOneAunt.cats || Sue.cats == -1 {
				if Sue.samoyeds == StarOneAunt.samoyeds || Sue.samoyeds == -1 {
					if Sue.pomeranians == StarOneAunt.pomeranians || Sue.pomeranians == -1 {
						if Sue.akitas == StarOneAunt.akitas || Sue.akitas == -1 {
							if Sue.vizslas == StarOneAunt.vizslas || Sue.vizslas == -1 {
								if Sue.goldfish == StarOneAunt.goldfish || Sue.goldfish == -1 {
									if Sue.trees == StarOneAunt.trees || Sue.trees == -1 {
										if Sue.cars == StarOneAunt.cars || Sue.cars == -1 {
											if Sue.perfumes == StarOneAunt.perfumes || Sue.perfumes == -1 {
												FoundAunts = append(FoundAunts, Sue)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//starTwo

		if Sue.children == StarOneAunt.children || Sue.children == -1 {
			if Sue.cats > StarOneAunt.cats || Sue.cats == -1 {
				if Sue.samoyeds == StarOneAunt.samoyeds || Sue.samoyeds == -1 {
					if Sue.pomeranians < StarOneAunt.pomeranians || Sue.pomeranians == -1 {
						if Sue.akitas == StarOneAunt.akitas || Sue.akitas == -1 {
							if Sue.vizslas == StarOneAunt.vizslas || Sue.vizslas == -1 {
								if Sue.goldfish < StarOneAunt.goldfish || Sue.goldfish == -1 {
									if Sue.trees > StarOneAunt.trees || Sue.trees == -1 {
										if Sue.cars == StarOneAunt.cars || Sue.cars == -1 {
											if Sue.perfumes == StarOneAunt.perfumes || Sue.perfumes == -1 {
												FoundAuntsStarTwo = append(FoundAuntsStarTwo, Sue)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

	}

	fmt.Printf("Found Aunts : %v ", FoundAuntsStarTwo)
	var starOne int = FoundAunts[0].number
	var starTwo int = FoundAuntsStarTwo[0].number

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}
