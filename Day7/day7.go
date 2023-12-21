package day7

import (
	"aoc-22/utils"
	"fmt"
	"strings"
)


var sample = "Day7/sample.txt"
var data = "Day7/data.txt"


func PartOne() {
    file := utils.OpenFile(sample)
    scanner := utils.GetFileLinesScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        if isCommand(line) {
            commandElements := strings.Split(line, " ")
            if commandElements[1] == "ls" {
                println("LIST")
            } else {

                /**
                * if '/' go to root
                * if '..' go to parent node
                * else 'x' go to x
                */

            }
        } else {
            output := []string{}
            output = append(output, line)

            for scanner.Scan() {
                line = scanner.Text()
                if !isCommand(line) {
                    output = append(output, line)
                } else {
                    // This currently skips the command after output, think about clean solution
                    break
                }
            }

            fmt.Printf("ALL OUTPUT: %v\n", output)

        }
    }
}

func PartTwo() {
}

func isCommand(line string) bool {
    return strings.HasPrefix(line, "$")

}
