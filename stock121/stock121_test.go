package stock121

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	prices []int
	output int
}

func TestStock121(t *testing.T){
	testcases := []TestCase{
		{prices: []int{7, 1, 5, 3, 6, 4},output: 5},
		{prices: []int{7,6,4,3,1},output: 0},
		{prices: []int{2,4,1},output: 2},
		{prices: []int{2,4,2,1},output: 2},
		{prices: []int{2,1,2,0,1},output: 1},
		{prices: []int{3,2,6,5,0,3},output: 4},
		{prices: []int{1,2},output: 1},
		{prices: []int{2,1,2,1,0,1,2},output: 2},
		{prices: []int{2,4,1,7},output: 6},
		{prices: []int{2,7,1,4},output: 5},
	}

	for i,testCase := range testcases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := Profit(testCase.prices)
			if !reflect.DeepEqual(output, testCase.output) {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})
	}
}