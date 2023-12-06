package day6

import (
	"aoc-22/utils"
	"bufio"
	"fmt"
	"os"
)

var sample = "Day6/sample.txt"
var data = "Day6/data.txt"

func PartOne() {
	file, err := os.Open(data)
	utils.CheckFile(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	items := []string{}
	marker := 0
	for i := 0; scanner.Scan(); i++ {
		if len(items) == 4 && !utils.ContainsDuplicate(items) {
			marker = i
			break
		}

		if len(items) > 3 {
			items = items[1:]
		}

		items = append(items, scanner.Text())
	}

	fmt.Println(marker)
}

func PartTwo() {

}
