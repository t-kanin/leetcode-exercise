package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for currIndex, currNumber := range nums {
		// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
		if requiredIndex, isPresent := m[target-currNumber]; isPresent {
			return []int{currIndex, requiredIndex}
		}
		m[currNumber] = currIndex
	}

	return []int{}
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9

	fmt.Print(twoSum(nums, target))
}
