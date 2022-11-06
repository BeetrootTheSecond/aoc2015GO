package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matchsticks struct {
	og            string
	inMemory      int
	stringLiteral int
}

func dayeight() {

	fmt.Println("Day Eight GO!")

	var data string = ReadFile("./data/day8/data.rawr")

	var lines []string = strings.Split(data, "\n")
	fmt.Printf("Data length : %v\n", len(lines))

	var results []Matchsticks
	totalInMemory := 0
	totalStringLiterals := 0
	totalencoded := 0
	for lineId := range lines {
		var line string = lines[lineId]
		totalInMemory += len(line)
		//StringLiterals := len(line)

		//first and last double quote removal
		s, _ := strconv.Unquote(line)

		e := strconv.Quote(line)

		fmt.Printf("%v\n", s)
		totalStringLiterals += len(s)
		fmt.Printf("%v\n", e)
		totalencoded += len(e)

		var newMatchstick Matchsticks = Matchsticks{line, len(line), len(s)}
		results = append(results, newMatchstick)
	}

	fmt.Printf("%v\n", results)
	fmt.Printf("InMemory  : %v\n", totalInMemory)
	fmt.Printf("totalStringLiterals  : %v\n", totalStringLiterals)
	fmt.Printf("totalencoded  : %v\n", totalencoded)

	var starOne = totalInMemory - totalStringLiterals
	var starTwo = totalencoded - totalInMemory

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}
