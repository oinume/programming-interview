package q

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayPairSum(t *testing.T) {
	result := ArrayPairSum([]int{8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7}, 10)
	fmt.Printf("result = %v\n", result)

	if !hasPair(result, []int{1, 9}) {
		t.Errorf("ArrayPairSum: result doesn't contain {1, 9}: %v", result)
	}
	if !hasPair(result, []int{3, 7}) {
		t.Errorf("ArrayPairSum: result doesn't contain {3, 7}: %v", result)
	}
}

func hasPair(values [][]int, pair []int) bool {
	for _, v := range values {
		if reflect.DeepEqual(pair, v) {
			return true
		}
	}
	return false
}
