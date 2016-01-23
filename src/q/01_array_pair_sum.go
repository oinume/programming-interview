package q

import (
	"fmt"
	"math"
	"sort"

	"golang.org/x/tools/container/intsets"
)

type IntPair [2]int

func ArrayPairSum(array []int, sum int) []IntPair {
	if len(array) < 2 {
		return nil
	}
	sort.Ints(array)

	result := make([]IntPair, 0, 100)
	left, right := 0, len(array) - 1
	currentSum := 0
	for ; left < right; {
		currentSum = array[left] + array[right]
		// fmt.Printf("currentSum = %v, left = %v, right = %v\n", currentSum, array[left], array[right])
		if currentSum == sum {
			result = append(result, IntPair{array[left], array[right]})
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

func ArrayPairSum2(array []int, sum int) []IntPair {
	if len(array) < 2 {
		return nil
	}

	seen := intsets.Sparse{}
	result := make([]IntPair, 0, 100)
	for _, n := range array {
		target := sum - n
		if !seen.Has(target) {
			fmt.Printf("Insert: target=%d, n=%d, seen=%v\n", target, n, seen.String())
			seen.Insert(n)
		} else {
			fmt.Printf("Has: target=%d, n=%d, seen=%v\n", target, n, seen.String())
			result = append(result, IntPair{
				int(math.Min(float64(n), float64(target))),
				int(math.Max(float64(n), float64(target))),
			})
		}
	}
	return result
}
