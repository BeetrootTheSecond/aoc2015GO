package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/valyala/fastjson"
)

func daytweleve() {

	fmt.Println("Day Twevele GO!")

	// b, err := os.ReadFile("./data/day12/data.json") // just pass the file name
	// if err != nil {
	// 	fmt.Print(err)
	// }

	var data string = ReadFile("./data/day12/data.json")

	var p fastjson.Parser
	v, err := p.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	// var objs []map[string]*json.RawMessage
	// if err := json.Unmarshal([]byte(b), &objs); err != nil {
	// 	panic(err)
	// }

	//v.Type()
	var valueArray = v.GetArray()

	var result = HandleArray(valueArray)

	fmt.Printf("Data : %v\n", result)

	var sum = 0
	for index := 0; index < len(result); index++ {
		sum += result[index]
	}

	var starOne = sum

	var starTwo = 0

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func HandleArray(valueArray []*fastjson.Value) []int {
	var FoundNumbers []int
	for i := 0; i < len(valueArray); i++ {
		var typeString = valueArray[i].Type()
		//fmt.Printf("Type : %v\n", typeString)
		switch typeString {
		case 1: //object
			{
				var thisValue = valueArray[i].GetObject()
				results := HandleObject(thisValue)
				fmt.Printf("found Object :  %v %v\n", thisValue, results)
				FoundNumbers = append(FoundNumbers, results...)
				break
			}
		case 2: //array
			{
				var thisValue = valueArray[i].GetArray()
				results := HandleArray(thisValue)
				//fmt.Printf("found Numbers : %v\n", results)
				FoundNumbers = append(FoundNumbers, results...)
				break
			}
		case 3: //string
			{
				//var thisValue = string(valueArray[i].GetStringBytes())

				//results := HandleArray(thisValue)
				//fmt.Printf("found string : %v\n", thisValue)
				//FoundNumbers = append(FoundNumbers, results...)
				break
			}
		case 4: //number
			{
				//var thisValue = valueArray[i].GetArray()
				//results := HandleArray(thisValue)
				//fmt.Printf("found Numbers : %v\n", valueArray[i].GetInt())
				FoundNumbers = append(FoundNumbers, valueArray[i].GetInt())
				break
			}

		default:
			{
				//panic("Found Unknown Type")
				fmt.Printf("Found Unknown Type %v \n", string(valueArray[i].Type()))
				break
			}

		}
	}

	return FoundNumbers

}

func HandleObject(valueObject *fastjson.Object) []int {
	var FoundNumbers []int
	var redFlag bool = false

	valueObject.Visit(func(k []byte, v *fastjson.Value) {
		if numberKey, err := strconv.Atoi(string(k)); err == nil {
			fmt.Printf("%q looks like a number.\n", numberKey)
			FoundNumbers = append(FoundNumbers, int(numberKey))
		}

		switch v.Type() {
		case 1: //object
			{
				var thisValue = v.GetObject()
				results := HandleObject(thisValue)
				fmt.Printf("found Object :  %v %v\n", thisValue, results)
				FoundNumbers = append(FoundNumbers, results...)
				break
			}
		case 2: //array
			{
				var thisValue = v.GetArray()
				results := HandleArray(thisValue)
				//fmt.Printf("found Numbers : %v\n", results)
				FoundNumbers = append(FoundNumbers, results...)
				break
			}
		case 3: //string
			{
				var thisValue = string(v.GetStringBytes())
				if thisValue == "red" {
					fmt.Printf("found string : %v\n", thisValue)
					redFlag = true
				}
				//results := HandleArray(thisValue)
				//fmt.Printf("found string : %v\n", thisValue)
				//FoundNumbers = append(FoundNumbers, results...)
				break
			}
		case 4: //number
			{
				//var thisValue = v.GetArray()
				//results := HandleArray(thisValue)
				//fmt.Printf("found Numbers : %v\n", v.GetInt())
				FoundNumbers = append(FoundNumbers, v.GetInt())
				break
			}

		default:
			{
				//panic("Found Unknown Type")
				fmt.Printf("Found Unknown Type %v \n", string(v.Type()))
				break
			}

		}
	})
	if !redFlag {
		return FoundNumbers
	}

	return make([]int, 0)
}
