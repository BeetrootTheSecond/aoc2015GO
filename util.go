package main

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) string {

	b, err := os.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'

	return str
}

func sum64(array []int64) int64 {
	var result int64 = 0
	for _, v := range array {
		result += v
	}
	return result
}
