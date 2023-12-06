package utils

import (
	"aoc-22/data_structures"
	"bufio"
	"fmt"
	"log"
	"os"
)

func CheckFile(err error) {
    Check(err, fmt.Sprintf("Error opening file: %v", err))
}


func Check(err error, msg string) {
	if err != nil {
		log.Printf(msg)
		panic(err)
	}
}

func GetFileLinesScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner
}

func GetMinimums(data []int64) (int, int64) {
	var amount int64
	var index int
	if len(data) > 0 {
		index = 0
		amount = data[index]
	}
	for idx, element := range data {
		if element < amount {
			index = idx
			amount = data[index]
		}
	}

	return index, amount
}

func Sum(data []int64) int64 {
	var total int64
	for _, e := range data {
		total += e
	}

	return total
}

func ContainsDuplicate(iterable []string) bool {
    seen := data_structures.NewStack()
    for _, element := range iterable {
        if seen.Contains(element) {
            return true
        } else {
            seen.Push(element)
        }
    }

    return false
}
