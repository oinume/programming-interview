package q

import (
	_ "fmt"
	_ "reflect"
	"testing"
)

func TestLargestContinuousSum(t *testing.T) {
	result1 := LargestContinousSum(100, 6, -46)
	if result1 != 106 {
		t.Errorf("LargestContinousSum: result is not 106: %v", result1)
	}

	result2 := LargestContinousSum(100, 6, -46, 10, 100)
	if result2 != 170 {
		t.Errorf("LargestContinousSum: result is not 170: %v", result2)
	}
}
