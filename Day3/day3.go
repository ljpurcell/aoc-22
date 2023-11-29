package day3

import (
	"aoc-22/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var data = "Day3/day3_data.txt"

func PartOne() {
	file, err := os.Open(data)
	utils.Check(err, "Could not open data file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters := strings.ToUpper(lowercaseLetters)

	total := 0
	for scanner.Scan() {
		items := scanner.Text()
		midPoint := len(items) / 2
		firstPack, secondPack := items[:midPoint], items[midPoint:]

		for _, item := range firstPack {
			if strings.Contains(secondPack, string(item)) {
				if unicode.IsUpper(item) {
					total += strings.Index(uppercaseLetters, string(item)) + 27
				} else {
					total += strings.Index(lowercaseLetters, string(item)) + 1
				}
				break
			}
		}
	}

	fmt.Println(total)
}

func PartTwo() {
	file, err := os.Open(data)
	utils.Check(err, "Could not open data file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters := strings.ToUpper(lowercaseLetters)

	total := 0
	packA, packB, packC := "", "", ""
	for scanner.Scan() {
		if packA == "" {
			packA = scanner.Text()
		} else if packB == "" {
			packB = scanner.Text()
		} else if packC == "" {
			packC = scanner.Text()
		}

		if packA != "" && packB != "" && packC != "" {
			for _, item := range packA {
				if strings.Contains(packB, string(item)) && strings.Contains(packC, string(item)) {
					if unicode.IsUpper(item) {
						total += strings.Index(uppercaseLetters, string(item)) + 27
					} else {
						total += strings.Index(lowercaseLetters, string(item)) + 1
					}
					break
				}
			}
			packA, packB, packC = "", "", ""
		}
	}
	fmt.Println(total)

}
