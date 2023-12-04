package day4

import (
	"aoc-22/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sample = "Day4/sample.txt"
var data = "Day4/data.txt"

func PartOne() {
	file, err := os.Open(data)
	utils.Check(err, "Could not open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")

		firstElf := strings.Split(elves[0], "-")
		secondElf := strings.Split(elves[1], "-")

        firstLower, _ := strconv.ParseInt(firstElf[0], 10, 64)
        secondLower, _ := strconv.ParseInt(secondElf[0], 10, 64)

        firstUpper, _ := strconv.ParseInt(firstElf[1], 10, 64)
        secondUpper, _ := strconv.ParseInt(secondElf[1], 10, 64)

		if firstLower <= secondLower && firstUpper >= secondUpper {
			count += 1
		} else if secondLower <= firstLower && secondUpper >= firstUpper {
			count += 1
		}
	}

    fmt.Println(count)
}
