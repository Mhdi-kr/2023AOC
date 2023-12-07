package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := int64(0)
	input := loadInputFromDisk()

	lines := strings.Split(input, "\n")

	for _, str := range lines {
		// runes := extractNumberCharactersFromString(str)

		// part 2
		runes := extractNumberCharactersOrNumberWordsFromString(str)

		result := makeNumberFromNumberCharacters(runes)
		sum += result
	}

	fmt.Println(sum) // correct answer: 54390
}

func loadInputFromDisk() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return string(content)
}

// func extractNumberCharactersFromString(input string) []rune {
// 	numbers := []rune{}

// 	// 47 - 57 => ASCII value of numbers 1 - 9
// 	for _, t := range input {
// 		if t > 46 && t < 58 {
// 			numbers = append(numbers, t)
// 		}
// 	}
// 	return numbers
// }

func extractNumberCharactersOrNumberWordsFromString(input string) []rune {
	numbers := []rune{}
	temp := input
	temp = normalizeNumberWords(input)

	// 47 - 57 => ASCII value of numbers 1 - 9
	for _, t := range temp {
		if t > 46 && t < 58 {
			numbers = append(numbers, t)
		}
	}

	return numbers
}

// func hasSubstringAt(source, wanted string) int {
// 	foundFirstIndex := -1
// 	hasFound := false
// 	sourceEnd := len(source) - 1
// 	wantedEnd := len(wanted) - 1

// 	for i := 0; i <= sourceEnd; i++ {

// 		if source[i] != wanted[0] {
// 			continue
// 		}

// 		if sourceEnd-i < wantedEnd {
// 			break
// 		}

// 		for j := 0; j <= wantedEnd; j++ {
// 			s := string(source[i+j])
// 			w := string(wanted[j])
// 			fmt.Println(s, w)

// 			if source[i+j] != wanted[j] {
// 				break
// 			}

// 			if j == wantedEnd {
// 				hasFound = true
// 			}
// 		}

// 		if hasFound {
// 			foundFirstIndex = i
// 			break
// 		}
// 	}

// 	return foundFirstIndex
// }

func normalizeNumberWords(input string) string {
	temp := input

	type NumberWord struct {
		value string
		at    int
	}
	numberWords := map[string]NumberWord{
		"one": {
			value: "1",
			at:    -1,
		},
		"two": {
			value: "2",
			at:    -1,
		},
		"three": {
			value: "3",
			at:    -1,
		},
		"four": {
			value: "4",
			at:    -1,
		},
		"five": {
			value: "5",
			at:    -1,
		},
		"six": {
			value: "6",
			at:    -1,
		},
		"seven": {
			value: "7",
			at:    -1,
		},
		"eight": {
			value: "8",
			at:    -1,
		},
		"nine": {
			value: "9",
			at:    -1,
		},
	}

	minimum := NumberWord{
		at:    -1,
		value: "",
	}

	for k, v := range numberWords {
		numberWords[k] = NumberWord{
			value: v.value,
			at:    strings.Index(temp, v.value),
		}
		if minimum.at < numberWords[k].at {
			minimum = numberWords[k]
		}
	}

	// strings.Replace(temp, minimum.value, minimum)

	return temp
}

func makeNumberFromNumberCharacters(input []rune) int64 {
	if len(input) == 0 {
		return 0
	}
	lastIndex := len(input) - 1
	first, _ := strconv.ParseInt(string(input[0]), 10, 64)
	if lastIndex == 0 {
		return (first * 10) + first
	}
	last, _ := strconv.ParseInt(string(input[lastIndex]), 10, 64)

	combined := (first * 10) + last
	return combined
}
