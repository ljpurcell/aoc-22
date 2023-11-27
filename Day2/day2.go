package day2

import (
	"aoc-22/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var data = "Day2/day2_data.txt"

func DayTwoPartOne() {
	file, err := os.Open(data)
	defer file.Close()
	utils.Check(err, "Could not open data file")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		signs := strings.Split(scanner.Text(), " ")
		opponent := signs[0]
		suggested := signs[1]

		if suggested == "X" {
			total += 1
			if opponent == "A" {
				total += 3
			} else if opponent == "C" {
				total += 6
			}
		} else if suggested == "Y" {
			total += 2
			if opponent == "B" {
				total += 3
			} else if opponent == "A" {
				total += 6
			}
		} else if suggested == "Z" {
			total += 3
			if opponent == "C" {
				total += 3
			} else if opponent == "B" {
				total += 6
			}
		}
	}

	fmt.Println(total)
}
