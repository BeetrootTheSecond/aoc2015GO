package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type reindeer struct {
	speed    int
	duration int
	rest     int
}

var reindeersStats map[string]reindeer = make(map[string]reindeer)

func dayfourteen() {

	fmt.Println("Day Fourteen GO!")

	var data string = ReadFile("./data/day14/data.rawr")
	var lines []string = strings.Split(data, "\n")
	fmt.Printf("%v \n", lines)

	var reindeers []string

	for lineId := range lines {
		line := lines[lineId]
		REG := regexp.MustCompile(`\d+`)
		exactedNumbers := REG.FindAllStringSubmatch(line, -1)

		words := strings.Split(line, " ")
		reindeersName := words[0]

		speed, _ := strconv.Atoi(exactedNumbers[0][0])
		duration, _ := strconv.Atoi(exactedNumbers[1][0])
		rest, _ := strconv.Atoi(exactedNumbers[2][0])
		nr := reindeer{speed, duration, rest}

		reindeersStats[reindeersName] = nr
		reindeers = append(reindeers, reindeersName)

		fmt.Printf("%v: %v \n", reindeersName, exactedNumbers)
	}

	fmt.Printf("reindeers %v \n", reindeersStats)

	_, starOneResults := Race(reindeers, 2503)
	starTwoResults := RaceNewRules(reindeers, 2503)

	sort.Slice(starOneResults[:], func(i, j int) bool {
		return starOneResults[i] > starOneResults[j]
	})

	sort.Slice(starTwoResults[:], func(i, j int) bool {
		return starTwoResults[i] > starTwoResults[j]
	})

	var starOne int = starOneResults[0]
	var starTwo int = starTwoResults[0]

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func Race(competitors []string, length int) (map[string]int, []int) {

	var results = make(map[string]int)
	var resultsInt []int
	for _, competitor := range competitors {

		//flags
		state := 0 // 0 for flying 1 for results[competitor].resting
		resting := -1
		flying := reindeersStats[competitor].duration
		flown := 0
		for i := 0; i < length; i++ {

			if resting == 0 && state == 1 {
				flying = reindeersStats[competitor].duration
				state = 0
			} else if flying == 0 && state == 0 {
				resting = reindeersStats[competitor].rest
				state = 1
			}

			if resting > 0 {
				resting--
			} else if flying > 0 {
				flying--
				flown += reindeersStats[competitor].speed
			}

		}
		results[competitor] = flown
		resultsInt = append(resultsInt, flown)
		fmt.Printf("Competitor : %v flow : %v\n", competitor, flown)
	}
	return results, resultsInt
}

func RaceNewRules(competitors []string, length int) []int {

	type raceDetails struct {
		state   int // 0 for flying 1 for resting
		resting int
		flying  int
		flown   int
		points  int
	}
	var results = make(map[string]raceDetails)
	var resultsInt []int

	for _, competitor := range competitors {
		rd := raceDetails{0, -1, reindeersStats[competitor].duration, 0, 0}
		results[competitor] = rd
	}

	for i := 0; i < length; i++ {
		for _, competitor := range competitors {
			if entry, ok := results[competitor]; ok {

				//flags

				if entry.resting == 0 && entry.state == 1 {
					entry.flying = reindeersStats[competitor].duration
					entry.state = 0
				} else if entry.flying == 0 && entry.state == 0 {
					entry.resting = reindeersStats[competitor].rest
					entry.state = 1
				}

				if entry.resting > 0 {
					entry.resting--
				} else if entry.flying > 0 {
					entry.flying--
					entry.flown += reindeersStats[competitor].speed
				}
				results[competitor] = entry
			}

		}

		var currentWinner []string
		var currentflown int = 0
		for _, competitor := range competitors {
			if results[competitor].flown > currentflown {
				currentflown = results[competitor].flown
				currentWinner = []string{competitor}
			} else if results[competitor].flown == currentflown {
				currentflown = results[competitor].flown
				currentWinner = append(currentWinner, competitor)
			}
		}

		for _, competitor := range currentWinner {
			if entry, ok := results[competitor]; ok {
				entry.points++
				results[competitor] = entry
			}
		}

	}

	for _, competitor := range competitors {

		resultsInt = append(resultsInt, results[competitor].points)
	}
	return resultsInt
}
