package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func loadInputFromDisk() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return string(content)
}

func findNumberOfWinningsInCurrent(current, winnings []int) int {
	found := 0
	for _, c := range current {
		for _, w := range winnings {
			if c == w {
				found += 1
				break
			}
		}
	}

	return found
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	content := loadInputFromDisk()
	lines := strings.Split(content, "\n")

	sum := 0

	for _, line := range lines {
		currentDirtyStr, winningsStr, _ := strings.Cut(line, " | ")
		_, currentStr, _ := strings.Cut(currentDirtyStr, ": ")
		currentStrArr := strings.Split(currentStr, " ")
		winningsStrArr := strings.Split(winningsStr, " ")

		currentArr := []int{}
		for _, i := range currentStrArr {
			if len(i) > 0 {
				res, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					panic("wrong number")
				}
				currentArr = append(currentArr, int(res))
			}
		}

		winningsArr := []int{}
		for _, i := range winningsStrArr {
			if len(i) > 0 {
				res, err := strconv.ParseInt(i, 10, 64)
				if err != nil {
					panic("wrong number")
				}
				winningsArr = append(winningsArr, int(res))
			}
		}

		n := findNumberOfWinningsInCurrent(currentArr, winningsArr)

		if n == 0 {
			continue
		}

		sum += powInt(2, n-1)
	}

	fmt.Println(sum) // part one correct answer: 23847
}
