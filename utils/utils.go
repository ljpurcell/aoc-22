package utils

import (
	"log"
)

func Check(err error, msg string) {
	if err != nil {
		log.Printf(msg)
		panic(err)
	}
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
