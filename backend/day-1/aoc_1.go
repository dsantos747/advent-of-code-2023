package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	tools "github.com/dsantos747/advent-of-code-2023/tools"
)

func numFormat(word string) string {
	wordMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":  "6",
		"seven":  "7",
		"eight":  "8",
		"nine":  "9",
	}

	if number, ok := wordMap[word]; ok {
		return number
	}
	return word
}

func part1(input []string) int {
	var calibSum int = 0
	re := regexp.MustCompile(`[0-9]`)

	for _,line := range input {
		matches := re.FindAll([]byte(line),-1)
		calib,_ := strconv.Atoi(string(matches[0]) + string(matches[len(matches)-1]))
		calibSum += calib
	}

	return calibSum
}

func part2(input []string) int {
	var calibSum int = 0
	re := regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)

	for _,line := range input {
		matches := re.FindAll([]byte(line),-1)
		first := numFormat(string(matches[0]))
		last := numFormat(string(matches[len(matches)-1]))
		calib,_ := strconv.Atoi(first + last)
		calibSum += calib
	}

	return calibSum
}

func main() {
	data,err := tools.ReadInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input := strings.Split(data, "\n")

	p1 := part1(input)
	fmt.Println("The answer to part 1 is",p1)

	p2 := part2(input)
	fmt.Println("The answer to part 2 is",p2)
}