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
