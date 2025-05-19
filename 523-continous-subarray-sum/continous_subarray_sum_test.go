package continoussubarraysum

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	Nums   []int
	K      int
	Output bool
}

func TestCheckSubarraySum(t *testing.T) {

	testCases := []TestCase{
		{
			Nums:   []int{23, 2, 4, 6, 7},
			K:      6,
			Output: true,
		},
		{
			Nums:   []int{23, 2, 6, 4, 7},
			K:      6,
			Output: true,
		},
		{
			Nums:   []int{23, 2, 6, 4, 7},
			K:      13,
			Output: false,
		},
		{
			Nums:   []int{23, 2, 4, 6, 6},
			K:      7,
			Output: true,
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase: %d\n", index+1), func(t *testing.T) {
			output := CheckSubarraySum(testCase.Nums,testCase.K)

			if !reflect.DeepEqual(output, testCase.Output) {
				t.Errorf("Expected %t | Output %t\n", testCase.Output, output)
			}

		})
	}
}
