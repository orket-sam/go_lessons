package main

import (
	"math"
	"slices"
)

func changeCalc(amount int) map[int]int {

	var result map[int]int
	var count int

	denominations := []int{1000, 500, 200, 100, 50, 40, 20, 10, 5, 1}
	if slices.Contains(denominations, amount) {
		return map[int]int{amount: 1}
	} else {
		for _, val := range denominations {

			if amount >= val {
				count = 0
				count = math.Abs(float64(amount) / val)

			}
		}
	}
	return result

}

func main() {
	result := changeCalc(100)
	for k, v := range result {
		println(k, v)
	}
}
