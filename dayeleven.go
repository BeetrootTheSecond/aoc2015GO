package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func dayeleven() {

	fmt.Println("Day Eleven GO!")

	var data string = ReadFile("./data/day11/data.rawr")

	//var lines []string = strings.Split(data, "\n")
	fmt.Printf("Data : %v\n", data)

	var splitPassword = []rune(data) //strings.Split(data, "")
	fmt.Printf("%v\n", splitPassword)
	var asciiArray []int

	for i := range splitPassword {
		fmt.Printf("%v\n", splitPassword[i])
		char := splitPassword[i]
		ascii := int(char)

		//fmt.Printf("%v , %v\n", char, ascii)
		//fmt.Printf("%c - %d\n", char, ascii)
		asciiArray = append(asciiArray, ascii)
	}

	// doubleLetter := pcre.MustCompile(`(.)\1{1,}`, 0)
	// matches := doubleLetter.MatcherString("aa bb", 0).Groups()
	// fmt.Printf("matches : %v\n", matches)

	result := nextVaildPassword(asciiArray)

	var resultString []string
	for resultI := 0; resultI < len(result); resultI++ {
		resultString = append(resultString, string(result[resultI]))
	}
	//testString := slices.Contains(asciiArray, 105)

	fmt.Printf("result: %v %v\n", data, strings.Join(resultString, ""))

	var starOne = strings.Join(resultString, "")

	resultStarTwo := nextVaildPassword(result)

	var resultStarTwoString []string
	for resultJ := 0; resultJ < len(resultStarTwo); resultJ++ {
		resultStarTwoString = append(resultStarTwoString, string(resultStarTwo[resultJ]))
	}
	var starTwo = strings.Join(resultStarTwoString, "")

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func nextVaildPassword(password []int) []int {

	var vaildPassword []int

	notVaild := false
	for !notVaild {
		vaildPassword = incrementPassword(password)

		var contains = !(slices.Contains(vaildPassword, 105) ||
			slices.Contains(vaildPassword, 111) ||
			slices.Contains(vaildPassword, 108))

		var doubleLetter = findMatchingPairs(vaildPassword)

		//straight of at least three
		var threeSeq = threeLetterSequence(vaildPassword)
		//doesn't contain i, o, or l
		if contains && doubleLetter && threeSeq {
			notVaild = true
		}

		//contains 2 different double letters

	}

	return vaildPassword
}

// ascii int a:97 - z:122
func incrementPassword(password []int) []int {
	for passwordChar := len(password) - 1; passwordChar >= 0; passwordChar-- {
		//add 1 to char
		password[passwordChar] = password[passwordChar] + 1
		//check if char if about z:122
		if password[passwordChar] == 123 {
			password[passwordChar] = 97
		} else {
			break
		}
	}
	return password
}

func findMatchingPairs(password []int) bool {
	var matchGroup []int

	for i := 1; i < len(password); i++ {
		if password[i-1] == password[i] {
			if !slices.Contains(matchGroup, password[i]) {
				matchGroup = append(matchGroup, password[i-1])
			}
		}
	}
	return len(matchGroup) >= 2
}

func threeLetterSequence(password []int) bool {
	var found bool = false

	for i := 2; i < len(password); i++ {
		if password[i-2] == password[i]-2 && password[i-1] == password[i]-1 {
			found = true
		}
	}
	return found

}
