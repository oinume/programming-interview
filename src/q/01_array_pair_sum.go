package q

import (
	"sort"
)

func ArrayPairSum(array []int, sum int) [][]int {
	if len(array) < 2 {
		return [][]int{}
	}
	sort.Ints(array)

	result := make([][]int, 0, 100)
	left, right := 0, len(array) - 1
	currentSum := 0
	for ; left < right; {
		currentSum = array[left] + array[right]
		// fmt.Printf("currentSum = %v, left = %v, right = %v\n", currentSum, array[left], array[right])
		if currentSum == sum {
			result = append(result, []int{array[left], array[right]})
			left += 1
			continue
		} else if currentSum < sum {
			left += 1 // search with bigger number
		} else {
			right -= 1 // search with smaller number
		}
	}
	return result
}
