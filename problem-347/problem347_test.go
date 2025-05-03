package problem347

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCases struct {
	Nums   []int
	K      int
	Output []int
}

func TestProblem347(t *testing.T) {
	testCases := []TestCases{
		{Nums: []int{1, 1, 1, 2, 2, 3}, K: 2, Output: []int{1, 2}},
		{Nums: []int{1}, K: 1, Output: []int{1}},
		{Nums: []int{1,1,2,2,3,3,3,4}, K: 2, Output: []int{2,3}},
	}
	for index, testCase := range testCases {

		t.Run(fmt.Sprintf("Test Case: %d\n", index+1), func(test *testing.T) {
			output := topKFrequent(testCase.Nums, testCase.K)
			if !reflect.DeepEqual(output, testCase.Output) {
				test.Errorf("Test Case: %d , Expected: %v, Got %v\n", index, testCase.Output, output)
			}
		})

	}
}
