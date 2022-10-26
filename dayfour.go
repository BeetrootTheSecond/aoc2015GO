package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func dayfour() {

	fmt.Println("Day Four GO!")

	var data string = ReadFile("./data/day4/data.rawr")

	// h := md5.New()
	// io.WriteString(h, "abcdef609043")
	// fmt.Printf("%x", h.Sum(nil))

	var currentInt int = 0
	var starOne int = 0
	var starTwo int = 0

	for {
		var currentStr = fmt.Sprintf("%v%v", data, currentInt) //strconv.FormatInt(int64(currentInt), 10)
		//fmt.Println(currentStr)
		//testStr
		//testStr := string(data + currentStr)
		//fmt.Println(testStr)

		//create md5
		h := md5.New()
		io.WriteString(h, currentStr)
		////io.WriteString(h, fmt.Sprintf("%v%b", data, starTwo))
		md5result := fmt.Sprintf("%x", h.Sum(nil)[:])
		//md5result := GetMD5Hash(testStr)

		var beginSub5 string = md5result[:5]
		var beginSub6 string = md5result[:6]
		//fmt.Printf("substring  5:  %v\n", beginSub5)
		//fmt.Printf("substring  6:  %v\n", beginSub6)
		if starOne == 0 && beginSub5 == "00000" {

			starOne = currentInt
			//fmt.Printf("Star one result : %v\n", starOne)
			//break

		}

		if starOne != 0 && beginSub6 == "000000" {
			starTwo = currentInt
			//fmt.Printf("Star two result : %v\n", starTwo)
			break
		}

		currentInt += 1
	}

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
