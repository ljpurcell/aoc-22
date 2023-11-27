package main

import (
	"aoc-22/utils"
	"bufio"
	"os"
	"strconv"
)

var data = "day1_data.txt"

func DayOnePartOne() {
	file, err := os.Open(data)
	defer file.Close()
	utils.Check(err, "Could not open day1_data.txt")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var maxElf, currentElf int64
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			amount, err := strconv.ParseInt(line, 10, strconv.IntSize)
			utils.Check(err, "Could not parse int")
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

func DayOnePartTwo() {
	file, err := os.Open(data)
	defer file.Close()
	utils.Check(err, "Could not open day1_data.txt")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	maxElves := []int64{0, 0, 0}
	var currentElf int64
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			amount, err := strconv.ParseInt(line, 10, 64)
			utils.Check(err, "Could not parse int")
			currentElf += amount
		} else {
			minIdx, minAmt := utils.GetMinimums(maxElves)
			if currentElf > minAmt {
				maxElves[minIdx] = currentElf
			}
			currentElf = 0
		}
	}

	sum := utils.Sum(maxElves)
	println(sum)
}
