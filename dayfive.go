package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func dayfive() {

	fmt.Println("Day Five GO!")

	var data string = ReadFile("./data/day5/data.rawr")

	var lines []string = strings.Split(data, "\n")

	//three vowels
	///[aeiou].*?[aeiou].*?[aeiou]/gm
	vowelReg, _ := regexp.Compile("[aeiou].*?[aeiou].*?[aeiou]")
	//fmt.Println(vowelReg.MatchString(lines[0]))
	//double letters
	//(.)\1
	regex := string(`(aa)|(bb)|(cc)|(dd)|(ee)|(ff)|(gg)|(hh)|(ii)|(jj)|(kk)|(ll)|(mm)|(nn)|(oo)|(pp)|(qq)|(rr)|(ss)|(tt)|(uu)|(vv)|(ww)|(xx)|(yy)|(zz)`)
	doubleLetterReg := regexp.MustCompile(regex)
	//fmt.Println(doubleLetterReg.MatchString(lines[0]))

	fmt.Println(!(strings.Contains(lines[3], "ab") ||
		strings.Contains(lines[3], "cd") ||
		strings.Contains(lines[3], "pq") ||
		strings.Contains(lines[3], "xy")))

	var filteredLines []string
	var filteredLines2 []string
	var filteredLines3 []string
	var filteredLines4 []string
	var filteredLines5 []string

	for i := range lines {

		if vowelReg.MatchString(lines[i]) {
			filteredLines = append(filteredLines, lines[i])

		}

	}

	for i := range filteredLines {

		if doubleLetterReg.MatchString(filteredLines[i]) {
			filteredLines2 = append(filteredLines2, filteredLines[i])

		}

	}

	for i := range filteredLines2 {

		if !(strings.Contains(filteredLines2[i], "ab") ||
			strings.Contains(filteredLines2[i], "cd") ||
			strings.Contains(filteredLines2[i], "pq") ||
			strings.Contains(filteredLines2[i], "xy")) {

			filteredLines3 = append(filteredLines3, filteredLines2[i])

		}

	}

	// fmt.Print(filteredLines)
	// fmt.Print(filteredLines2)
	// fmt.Print(filteredLines3)

	patterntwiceReg := pcre.MustCompile(`([\w][\w]).*\1`, 0)
	patternonceReg := pcre.MustCompile(`([\w]).\1`, 0)

	for i := range lines {
		matches := patterntwiceReg.MatcherString((lines[i]), 0).Matches()
		matches2 := patternonceReg.MatcherString((lines[i]), 0).Matches()
		if matches && matches2 {

			filteredLines4 = append(filteredLines4, lines[i])

		}

	}

	fmt.Print(len(filteredLines4))
	fmt.Print(len(filteredLines5))

	var starOne = len(filteredLines3[:])
	var starTwo = len(filteredLines4[:])

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}
