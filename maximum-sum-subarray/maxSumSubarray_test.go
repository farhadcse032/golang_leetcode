package maxsumsubarray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	Input    []int
	Expected []int
}

func TestMaxSumSubArray(t *testing.T) {
	testcases := []TestCase{
		{
			Input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			Expected: []int{4, -1, 2, 1},
		},
		{
			Input:    []int{1},
			Expected: []int{1},
		},
		{
			Input:    []int{5, 4, -1, 7, 8},
			Expected: []int{5, 4, -1, 7, 8},
		},
	}

	for index, testcase := range testcases {
		t.Run(fmt.Sprintf("test number:%d\n", index+1), func(t *testing.T) {
			output := MaxSumSubArrayV2(testcase.Input)
			if !reflect.DeepEqual(output, testcase.Expected) {
				t.Errorf("Expected %d, Got %d", testcase.Expected, output)
			}
		})
	}
}
