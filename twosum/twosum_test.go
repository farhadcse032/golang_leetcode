package twosum


import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	output []int
}

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{nums: []int{2, 7, 11, 15}, target: 9, output: []int{2, 7}},
		{nums: []int{3, 2, 4}, target: 6, output: []int{2, 4}},
		{nums: []int{3, 2, 4}, target: 8, output: []int{-1, -1}},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := TwoSum(testCase.nums, testCase.target)
			if !reflect.DeepEqual(output, testCase.output) {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}

func TestTwoSumIndex(t *testing.T) {
	testCases := []TestCase{
		{nums: []int{2, 7, 11, 15}, target: 9, output: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, output: []int{1, 2}},
		{nums: []int{3, 2, 4}, target: 8, output: []int{-1, -1}},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := TwoSumIndex(testCase.nums, testCase.target)
			if !reflect.DeepEqual(output, testCase.output) {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}else{
				fmt.Printf("Passed Iteration %d\n",(i+1))
			}
		})

	}
}