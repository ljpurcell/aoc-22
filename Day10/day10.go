package day10

import (
	"aoc-22/utils"
	"fmt"
	"strconv"
	"strings"
)

const sample = "Day10/sample.txt"
const data = "Day10/data.txt"

func PartOne() {
	file := utils.OpenFile(data)

	scanner := utils.GetFileLinesScanner(file)

	sum := 0
	register := 1

	for cycles, nextRelevantCycle := 1, 20; scanner.Scan(); cycles++ {
		if cycles == nextRelevantCycle {
			nextRelevantCycle += 40
			sum += register * cycles
		}

        instruction := strings.Split(scanner.Text(), " ")
		if len(instruction) == 2 {
			cycles++

			if cycles == nextRelevantCycle {
				nextRelevantCycle += 40
				sum += register * cycles
			}

			val, err := strconv.ParseInt(instruction[1], 10, strconv.IntSize)
			utils.Check(err, "Could not parse integer value")
			register += int(val)
		}
	}

    fmt.Println(sum)
}

func PartTwo() {}
