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
		{prices: []int{2,4,1},output: 2},
		{prices: []int{2,4,2,1},output: 2},
		{prices: []int{2,1,2,0,1},output: 1},
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