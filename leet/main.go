package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9

	twoSum(nums, target)
	addNodeList([]int{2, 4, 3}, []int{5, 6, 4})
}

func twoSum(nums []int, target int) {

	for i, a := range nums {
		for j, b := range nums {
			if a+b == target {
				output := [2]int{i, j}
				fmt.Println(output)
				return
			}
		}
	}
}

func addNodeList(a, b []int) {

	slices.Reverse(a)
	slices.Reverse(b)

	var aString string
	var bString string

	for _, val := range a {
		aString += strconv.Itoa(val)
	}
	for _, val := range b {
		bString += strconv.Itoa(val)
	}
	aInt, _ := strconv.Atoi(aString)
	bInt, _ := strconv.Atoi(bString)

	fmt.Println(aInt + bInt)

}
