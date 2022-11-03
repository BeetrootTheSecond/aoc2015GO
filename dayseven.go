package main

import (
	"fmt"
	"strconv"
	"strings"
)

type wireLayout struct {
	result   uint16
	assigned bool
}

func dayseven() {

	fmt.Println("Day Seven GO!")

	var data string = ReadFile("./data/day7/data.rawr")

	var lines []string = strings.Split(data, "\n")
	fmt.Printf("Data length : %v\n", len(lines))

	var wireResults = make(map[string]wireLayout)

	fmt.Printf("wireResults length : %v\n", len(wireResults))

	wireResults = processWires(wireResults, lines)

	var wireResultsStar2 = make(map[string]wireLayout)
	wireResultsStar2["b"] = wireLayout{wireResults["a"].result, true}

	wireResultsStar2 = processWires(wireResultsStar2, lines)
	// // sample test
	// fmt.Printf("d : %v\n", wireResults["d"].result)
	// fmt.Printf("e : %v\n", wireResults["e"].result)
	// fmt.Printf("f : %v\n", wireResults["f"].result)
	// fmt.Printf("g : %v\n", wireResults["g"].result)
	// fmt.Printf("h : %v\n", wireResults["h"].result)
	// fmt.Printf("i : %v\n", wireResults["i"].result)
	// fmt.Printf("x : %v\n", wireResults["x"].result)
	// fmt.Printf("y : %v\n", wireResults["y"].result)

	var starOne = wireResults["a"].result
	var starTwo = wireResultsStar2["a"].result

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func processWires(wireResults map[string]wireLayout, lines []string) map[string]wireLayout {

	for {
		var loopValues []string

		for operationIndex := range lines {
			//fmt.Printf("Data length : %v\n", lines[operationIndex])

			var splitOperation = strings.Split(lines[operationIndex], " -> ")
			var identifier = splitOperation[1]
			//fmt.Printf("Data identifier : %v\n", identifier)

			switch {
			case strings.Contains(splitOperation[0], "AND"):
				{

					splitOperationRemians := strings.Split(splitOperation[0], " ")
					var firstParameterCheck = wireResults[splitOperationRemians[0]].assigned
					var firstParameter = wireResults[splitOperationRemians[0]].result

					if resultfirst, err := strconv.Atoi(splitOperationRemians[0]); err == nil {
						//fmt.Printf("%q looks like a number.\n", splitOperationRemians[0])
						firstParameterCheck = true
						firstParameter = uint16(resultfirst)
					}

					if firstParameterCheck && wireResults[splitOperationRemians[2]].assigned {

						result := firstParameter & wireResults[splitOperationRemians[2]].result

						wireResults[identifier] = wireLayout{result, true}

						//fmt.Printf("AND : %v\n", wireResults[identifier])
						fmt.Printf("AND %v  :  %v\n", lines[operationIndex], wireResults[identifier].result)
					} else {
						loopValues = append(loopValues, lines[operationIndex])
						//fmt.Print("push-AND")
					}
				}
			case strings.Contains(splitOperation[0], "OR"):
				{

					splitOperationRemians := strings.Split(splitOperation[0], " ")

					if wireResults[splitOperationRemians[0]].assigned && wireResults[splitOperationRemians[2]].assigned {

						result := wireResults[splitOperationRemians[0]].result | wireResults[splitOperationRemians[2]].result

						wireResults[identifier] = wireLayout{result, true}

						////fmt.Printf("OR: %v\n", wireResults[identifier])
						fmt.Printf("OR %v  :  %v\n", lines[operationIndex], wireResults[identifier].result)
					} else {
						loopValues = append(loopValues, lines[operationIndex])
						//fmt.Print("push-OR")
					}
				}
			case strings.Contains(splitOperation[0], "LSHIFT"):
				{
					splitOperationRemians := strings.Split(splitOperation[0], " ")
					if wireResults[splitOperationRemians[0]].assigned {

						result := wireResults[splitOperationRemians[0]].result
						ocours, _ := strconv.Atoi(splitOperationRemians[2])
						//for i := 0; i < ocours; i++ {
						result = result << ocours
						//}

						wireResults[identifier] = wireLayout{result, true}

						////fmt.Printf("LSHIFT: %v\n", wireResults[identifier])
						fmt.Printf("LSHIFT %v  :  %v\n", lines[operationIndex], wireResults[identifier].result)
					} else {
						loopValues = append(loopValues, lines[operationIndex])
						//fmt.Print("push-L")
					}
				}
			case strings.Contains(splitOperation[0], "RSHIFT"):
				{
					splitOperationRemians := strings.Split(splitOperation[0], " ")

					if wireResults[splitOperationRemians[0]].assigned {

						result := wireResults[splitOperationRemians[0]].result
						ocours, _ := strconv.Atoi(splitOperationRemians[2])
						//for i := 0; i < ocours; i++ {
						result = result >> ocours
						//}

						wireResults[identifier] = wireLayout{result, true}

						////fmt.Printf("RSHIFT: %v\n", wireResults[identifier])
						fmt.Printf("RSHIFT %v  :  %v\n", lines[operationIndex], wireResults[identifier].result)
					} else {
						loopValues = append(loopValues, lines[operationIndex])
						//fmt.Print("push-R")
					}
				}
			case strings.Contains(splitOperation[0], "NOT"):
				{
					splitOperationRemians := strings.Split(splitOperation[0], "NOT ")
					if wireResults[splitOperationRemians[1]].assigned {

						//fmt.Print(splitOperationRemians[1])
						result := ^uint16(wireResults[splitOperationRemians[1]].result)
						//fmt.Print(result)
						wireResults[identifier] = wireLayout{result, true}
						//fmt.Printf("NOT: %v\n", wireResults[identifier])
						fmt.Printf("NOT %v  :  %v\n", lines[operationIndex], wireResults[identifier].result)
					} else {
						loopValues = append(loopValues, lines[operationIndex])
						//fmt.Print("push-NOT")
					}
				}
			default:
				{
					if !wireResults[identifier].assigned {

						if resultfirst, err := strconv.Atoi(splitOperation[0]); err == nil {
							//fmt.Printf("%q looks like a number.\n", splitOperationRemians[0])

							wireResults[identifier] = wireLayout{uint16(resultfirst), true}
							//fmt.Printf("assignment : %v\n", wireResults[identifier])
							fmt.Printf("assignment Direct %v  :  %v\n", lines[operationIndex], wireResults[identifier].result)
						} else if wireResults[splitOperation[0]].assigned {
							wireResults[identifier] = wireLayout{wireResults[splitOperation[0]].result, true}
							fmt.Printf("assignment %v  :  %v\n", lines[operationIndex], wireResults[identifier].result)
						} else {
							loopValues = append(loopValues, lines[operationIndex])
						}

					}
				}

			}

		}

		lines = loopValues
		fmt.Printf("%v\n", len(loopValues))
		if len(lines) == 0 {
			break
		}
	}

	return wireResults
}
