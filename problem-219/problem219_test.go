package problem219

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCases struct {
	Input  []int
	Given  int
	Output bool
}

func TestProblem219(t *testing.T) {
	testcases := []TestCases{
		// {Input: []int{1, 2, 3, 1}, Given: 3, Output: true},
		// {Input: []int{1, 0, 1, 1}, Given: 1, Output: true},
		// {Input: []int{1, 2, 3, 1, 2, 3}, Given: 2, Output: false},
		{Input: []int{1, 2, 3,4, 1, 2, 5,1}, Given: 2, Output: true},
	}

	for i, testCase := range testcases {
		t.Run(fmt.Sprintf("Test Case : %d\n", (i+1)), func(test *testing.T) {
			output := Problem219(testCase.Input, testCase.Given)
			if !reflect.DeepEqual(output, testCase.Output) {
				test.Errorf("Expected : %+v, Output : %+v", testCase.Output, output)
			}
		})
	}
}
