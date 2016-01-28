package q

import (
	_ "fmt"
)

func LargestContinousSum(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	var currentSum, sum int
	for _, v := range values {
		currentSum += v
		if currentSum > sum {
			sum = currentSum
		}
	}
	return sum
}
