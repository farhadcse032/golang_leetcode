package removeduplicatefromsortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	Input  []int
	Output int
}

func TestRemoveDuplicate(t *testing.T) {

	testCases := []TestCase{

		{
			Input:  []int{1, 1, 2},
			Output: 2,
		},
		{
			Input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Output: 5,
		}, {
			Input:  []int{1, 1, 1, 2},
			Output: 2,
		}, {
			Input:  []int{1},
			Output: 1,
		},
	}

	for index, v := range testCases {
		t.Run(fmt.Sprintf("TestCase: %d\n", index+1), func(t *testing.T) {
			output := RemoveDuplicate(v.Input)

			if !reflect.DeepEqual(output, v.Output) {
				t.Errorf("Expected: %d | Got: %d\n", v.Output, output)
			}
		})
	}
}
