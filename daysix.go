package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	newState   int
	startingX  int
	startingY  int
	endingX    int
	endingY    int
	brightness int
}

func daysix() {

	fmt.Println("Day Six GO!")

	var data string = ReadFile("./data/day6/sample.rawr")

	var lines []string = strings.Split(data, "\n")
	fmt.Printf("Data length : %v\n", len(lines))

	var processLines []instruction

	//parse data into readable format
	for i := range lines {
		var line = lines[i]
		var newInstruction = instruction{}
		//find start
		if strings.Contains(line, "turn on ") {
			newInstruction.newState = 1
			line = strings.Replace(line, "turn on ", "", 1)
		} else if strings.Contains(line, "turn off ") {
			newInstruction.newState = 0
			line = strings.Replace(line, "turn off ", "", 1)
		} else if strings.Contains(line, "toggle ") {
			newInstruction.newState = 2
			line = strings.Replace(line, "toggle ", "", 1)
		}

		var pointRanges []string = strings.Split(line, " through ")

		var StartingRange = strings.Split(pointRanges[0], ",")

		startingX, errStartingX := strconv.Atoi(StartingRange[0])
		startingY, errStartingY := strconv.Atoi(StartingRange[1])

		newInstruction.startingX = startingX
		newInstruction.startingY = startingY

		if errStartingX != nil {
			// ... handle error
			panic(errStartingX)
		}
		if errStartingY != nil {
			// ... handle error
			panic(errStartingY)
		}

		var endingRange = strings.Split(pointRanges[1], ",")
		endingX, errEndingX := strconv.Atoi(endingRange[0])
		endingY, errEndingY := strconv.Atoi(endingRange[1])

		newInstruction.endingX = endingX
		newInstruction.endingY = endingY

		if errEndingX != nil {
			// ... handle error
			panic(errEndingX)
		}
		if errEndingY != nil {
			// ... handle error
			panic(errEndingY)
		}

		//star two
		newInstruction.brightness = 0
		//fmt.Printf("Data length : %v remaining line : %v \n", newInstruction, line)

		processLines = append(processLines, newInstruction)
	}

	//fmt.Printf("processLines : %v\n", processLines)
	//strings.Contains(lines[3], "cd")

	var lightsGrid [1000][1000]bool
	fmt.Printf("Data length : %v\n", len(lightsGrid))

	var lightsGridBrightness [1000][1000]int
	var frames []*image.Paletted
	var delays []int
	delays = append(delays, 30)

	for i := range processLines {
		var currentInstruction = processLines[i]
		fmt.Printf("currentInstruction : %v \n", currentInstruction)

		for x := currentInstruction.startingX; x <= currentInstruction.endingX; x++ {
			for y := currentInstruction.startingY; y <= currentInstruction.endingY; y++ {
				switch currentInstruction.newState {
				case 0: // turn off
					{
						lightsGrid[x][y] = false
						lightsGridBrightness[x][y] -= 1
						if lightsGridBrightness[x][y] < 0 {
							lightsGridBrightness[x][y] = 0
						}

						break

					}
				case 1: // turn on
					{
						lightsGrid[x][y] = true
						lightsGridBrightness[x][y] += 1
						break
					}
				case 2: //toggle
					{
						lightsGrid[x][y] = !lightsGrid[x][y]
						lightsGridBrightness[x][y] += 2
						break
					}
				default:
					{
						panic("Well Shit")
					}
				}

				frame := createImage(lightsGrid)
				frames = append(frames, frame)
				//delays = append(delays, 5)

			}

		}
		// frame := createImage(lightsGrid)
		// frames = append(frames, frame)
		// delays = append(delays, 10)
	}
	//gif creation

	f, _ := os.OpenFile("data/day6/DataAnimatedAllTheFrames.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: frames,
		Delay: delays,
	})

	var countLitLights int = 0
	var countLitbrightness int = 0

	//create image ?
	width := 1000
	height := 1000
	// upLeft := image.Point{0, 0}
	// lowRight := image.Point{width, height}

	// img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// red := color.RGBA{255, 0, 0, 0xff}
	// green := color.RGBA{0, 255, 0, 0xff}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if lightsGrid[x][y] {
				//img.Set(x, y, green)
				countLitLights = countLitLights + 1
				countLitbrightness += lightsGridBrightness[x][y]
			} else {
				//img.Set(x, y, red)
				countLitbrightness += lightsGridBrightness[x][y]
			}
		}
	}

	// // Encode as PNG.
	// f, _ := os.Create("data/day6/Data.png")
	// png.Encode(f, img)

	var starOne = countLitLights
	var starTwo = countLitbrightness

	fmt.Printf("Star one result : %v\n", starOne)

	fmt.Printf("Star two result : %v\n", starTwo)

}

func createImage(lightsGrid [1000][1000]bool) *image.Paletted {
	width := 1000
	height := 1000
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	white := color.RGBA{255, 255, 255, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}

	img := image.NewPaletted(image.Rectangle{upLeft, lowRight}, palette.WebSafe)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if lightsGrid[x][y] {
				img.Set(x, y, white)
			} else {
				img.Set(x, y, black)

			}
		}
		fmt.Printf("row : %v\n", x)
	}

	return img
}
