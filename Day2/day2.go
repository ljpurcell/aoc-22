package day2

import (
	"aoc-22/utils"
	"bufio"
	"fmt"
	"os"
)

var data = "Day2/day2_data.txt"

func DayTwoPartOne() {
	file, err := os.Open(data)
	defer file.Close()
	utils.Check(err, "Could not open data file")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

    // A : X - Rock
    // B : Y - Paper
    // C : Z - Scissors

    scores := map[string]int{
        "B X": 1,
        "C Y": 2,
        "A Z": 3,
        "A X": 4,
        "B Y": 5,
        "C Z": 6,
        "C X": 7,
        "A Y": 8,
        "B Z": 9,
    }

	total := 0
	for scanner.Scan() {
        total += scores[scanner.Text()]
	}

	fmt.Println(total)
}
