package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var ingredients []ingredient

func dayfifteen() {

	fmt.Println("Day Fifteen GO!")

	var data string = ReadFile("./data/day15/data.rawr")
	var lines []string = strings.Split(data, "\n")
	fmt.Printf("%v \n", lines)

	for _, line := range lines {
		//fmt.Printf("%v \n", line)
		var ingredientName = strings.Split(line, ":")[0]
		REG := regexp.MustCompile(`-?\d+`)
		exactedNumbers := REG.FindAllStringSubmatch(line, -1)

		capacity, _ := strconv.Atoi(exactedNumbers[0][0])
		durability, _ := strconv.Atoi(exactedNumbers[1][0])
		flavor, _ := strconv.Atoi(exactedNumbers[2][0])
		texture, _ := strconv.Atoi(exactedNumbers[3][0])
		calories, _ := strconv.Atoi(exactedNumbers[4][0])

		var ni = ingredient{ingredientName, capacity, durability, flavor, texture, calories}

		ingredients = append(ingredients, ni)
	}

	fmt.Printf("%v \n", ingredients)

	allPosibities := teaspoons(100, len(ingredients))

	var starOneResults []int
	var starTwoResults []int
	for _, teaspoonsPossibity := range allPosibities {

		//fmt.Printf("%v \n", teaspoonsPossibity)
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0
		calories := 0

		for i := 0; i < len(teaspoonsPossibity); i++ {
			capacity += ingredients[i].capacity * teaspoonsPossibity[i]
			durability += ingredients[i].durability * teaspoonsPossibity[i]
			flavor += ingredients[i].flavor * teaspoonsPossibity[i]
			texture += ingredients[i].texture * teaspoonsPossibity[i]
			calories += ingredients[i].calories * teaspoonsPossibity[i]
		}

		if capacity < 0 {
			capacity = 0
		}
		if durability < 0 {
			durability = 0
		}
		if flavor < 0 {
			flavor = 0
		}
		if texture < 0 {
			texture = 0
		}

		total := capacity * durability * flavor * texture
		fmt.Printf("%v \n", total)
		starOneResults = append(starOneResults, total)

		if calories == 500 {
			starTwoResults = append(starTwoResults, total)
		}
	}

	sort.Slice(starOneResults[:], func(i, j int) bool {
		return starOneResults[i] > starOneResults[j]
	})
	sort.Slice(starTwoResults[:], func(i, j int) bool {
		return starTwoResults[i] > starTwoResults[j]
	})

	var starOne int = starOneResults[0] //19813200 is too low //18734912
	var starTwo int = starTwoResults[0]

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

// find all options for teaspoons
func teaspoons(teaspoonsLimit int, splitsBetween int) [][]int {

	//positions
	//var a, b, c, d = 0, 0, 0, 0
	var allPosibities [][]int

	for a := range makeRange(0, 100) {
		for b := range makeRange(a+1, 100) {
			if 100-(a+b) <= b {
				break
			}
			for c := range makeRange(b+1, 100) {
				if 100-(a+b+c) <= c {
					break
				}
				d := 100 - (a + b + c)
				position := []int{a, b, c, d}
				//fmt.Printf(" %v", position)
				allPosibities = append(allPosibities, permutationsInt(position)...)
			}
		}

	}

	fmt.Printf("%v \n", len(allPosibities))

	return allPosibities

}

func removeDuplicateValues(intSlice [][]int) [][]int {
	keys := make(map[string]bool)
	list := [][]int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		entryStr := strings.Join(strings.Fields(fmt.Sprint(entry)), ",")
		if _, value := keys[string(entryStr)]; !value {
			keys[string(entryStr)] = true
			list = append(list, entry)
		}
	}
	return list
}

func CheckIfMatch(intSlice []int) bool {

	//get last element
	lastElement := intSlice[len(intSlice)-1]

	for i := len(intSlice) - 2; i >= 0; i-- {
		if intSlice[i] != lastElement {
			return false
		}
	}
	return true
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
