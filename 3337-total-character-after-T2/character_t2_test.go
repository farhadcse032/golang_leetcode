package totalcharacteraftert2

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	S      string
	T      int
	nums   []int
	Output int
}

func TestCharacterT2(t *testing.T) {

	testCases := []TestCase{
		{
			S:      "abcyy",
			T:      2,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
			Output: 7,
		},
		{
			S:      "azbk",
			T:      1,
			nums:   []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			Output: 8,
		},
	}

	for index, v := range testCases {

		t.Run(fmt.Sprintf("Test Case: %d\n", index+1), func(t *testing.T) {
			output := LengthAfterTransformations(v.S, v.T, v.nums)

			if !reflect.DeepEqual(output, v.Output) {
				t.Errorf("Expected: %d | Got %d \n", v.Output, output)
			}
		})
	}
}
