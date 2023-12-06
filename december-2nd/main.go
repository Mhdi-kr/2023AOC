package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := loadInputFromDisk()
	lines := strings.Split(content, "\n")

	type Game struct {
		Id     int
		Blues  []int
		Greens []int
		Reds   []int
	}

	games := make([]Game, 0)

	for _, line := range lines {
		game := Game{}
		gameStr, colorsStr, _ := strings.Cut(line, ": ")
		_, gameIDStr, _ := strings.Cut(gameStr, " ")

		gameID, err := strconv.ParseInt(gameIDStr, 10, 64)
		if err != nil {
			panic("wrong id")
		}

		game.Id = int(gameID)

		gameSets := strings.Split(colorsStr, "; ")

		for _, set := range gameSets {
			colorsArr := strings.Split(set, ", ")
			for _, colorStr := range colorsArr {
				valueStr, _, _ := strings.Cut(colorStr, " ")
				colorValue, err := strconv.ParseInt(valueStr, 10, 64)
				if err != nil {
					panic("wrong value")
				}

				if strings.Contains(colorStr, "red") {
					game.Reds = append(game.Reds, int(colorValue))
				}
				if strings.Contains(colorStr, "green") {
					game.Greens = append(game.Greens, int(colorValue))
				}
				if strings.Contains(colorStr, "blue") {
					game.Blues = append(game.Blues, int(colorValue))
				}
			}
		}
		games = append(games, game)
	}

	fmt.Println(games)

	sum := 0
	for _, game := range games {

		isBlueOK := true
		isRedOK := true
		isGreenOk := true

		for _, r := range game.Reds {
			if r > 12 {
				isRedOK = false
			}
		}
		for _, g := range game.Greens {
			if g > 13 {
				isGreenOk = false
			}
		}
		for _, b := range game.Blues {
			if b > 14 {
				isBlueOK = false
			}
		}

		if isBlueOK && isRedOK && isGreenOk {
			sum += game.Id
		}
	}

	// part 2
	sum2 := 0
	for _, game := range games {
		maxBlue := 0
		maxRed := 0
		maxGreen := 0

		for _, r := range game.Reds {
			if r > maxRed {
				maxRed = r
			}
		}
		for _, g := range game.Greens {
			if g > maxGreen {
				maxGreen = g
			}
		}
		for _, b := range game.Blues {
			if b > maxBlue {
				maxBlue = b
			}
		}

		sum2 += maxBlue * maxGreen * maxRed
	}

	println(sum2)

}

func loadInputFromDisk() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return string(content)
}
