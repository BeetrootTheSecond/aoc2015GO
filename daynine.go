package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type route struct {
	distance int
	pointOne string
	pointTwo string
}

func daynine() {

	fmt.Println("Day Nine GO!")

	var data string = ReadFile("./data/day9/data.rawr")

	locationsCheck := make(map[string]bool)
	locations := []string{}

	var lines []string = strings.Split(data, "\n")
	fmt.Printf("Data length : %v\n", len(lines))
	var routes []route
	for routeId := range lines {
		var routestr = lines[routeId]
		routeSplit := strings.Split(routestr, " = ")
		distance, _ := strconv.Atoi(string(routeSplit[1]))
		points := strings.Split(routeSplit[0], " to ")

		routes = append(routes, route{distance, points[0], points[1]})
		routes = append(routes, route{distance, points[1], points[0]})

		if !locationsCheck[points[0]] {
			locations = append(locations, points[0])
			locationsCheck[points[0]] = true
		}
		if !locationsCheck[points[1]] {
			locations = append(locations, points[1])
			locationsCheck[points[1]] = true
		}
	}
	fmt.Printf("Routes : %v \n ", routes)
	fmt.Printf("Locations : %v \n ", locations)

	// for locationID := range locations[:] {
	// 	fmt.Printf("locationsB : %v\n", locations)
	// 	startingLocation := locations[locationID]
	// 	fmt.Printf("Starting Location : %v\n", startingLocation)

	// }
	var distances []int
	var permuationsArr = permutations(locations)
	fmt.Printf("permuationsArr : %v\n", permuationsArr)

	for per := range permuationsArr {
		var journey = permuationsArr[per]
		var distance int = 0
		for stop := range journey {
			var stopName = journey[stop]
			if stop < len(journey)-1 {
				var nextStop = journey[stop+1]
				filtered := []int{}
				for i := range routes {
					var route = routes[i]
					if route.pointOne == stopName && route.pointTwo == nextStop {
						filtered = append(filtered, route.distance)
					}
				}

				if len(filtered) == 1 {
					distance += filtered[0]
				} else {
					panic("more then 1 found ")
				}

				//fmt.Printf("stopName : %v, Next Stop : %v, Disatance : %v\n", stopName, nextStop, distances)

			}
		}
		distances = append(distances, distance)
	}

	//fmt.Printf("distances : %v\n", distances)
	sort.Ints(distances[:])
	//fmt.Printf("distances : %v\n", distances)
	// find shortest

	var starOne = distances[0]
	var starTwo = distances[len(distances)-1]

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	result := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return result
}
