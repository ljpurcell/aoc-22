package main

import (
	"aoc-22/utils"
	"bufio"
	"os"
	"strconv"
)

func DayOnePartOne() {
	file, err := os.Open("day1_data.txt")
	defer file.Close()
	utils.CheckAndLog(err, "Could not open day1_data.txt")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var maxElf, currentElf int64
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			amount, err := strconv.ParseInt(line, 10, strconv.IntSize)
			utils.CheckAndLog(err, "Could not parse int")
			currentElf += amount
		} else {
			if currentElf > maxElf {
				maxElf = currentElf
			}
			currentElf = 0
		}
	}

    println(maxElf)
}
