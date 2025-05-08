package palindromicsubstrings

import (
	"fmt"
	"reflect"
	"testing"
)

type TeseCase struct {
	Input    string
	Expected int
}

func TestPalindromicsubstrings(t *testing.T) {

	testCases := []TeseCase{
		{
			Input:    "abc",
			Expected: 3,
		},
		{
			Input:    "aaa",
			Expected: 6,
		},
		{
			Input:    "aba",
			Expected: 4,
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase: %d", index+1), func(t *testing.T) {
			output := Palindromicsubstrings(testCase.Input)

			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Errorf("Expected %d | Output %d\n", testCase.Expected, output)
			}

		})
	}
}
