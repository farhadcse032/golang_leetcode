package ransom

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	String1 string
	String2 string
	Output  string
}

func TestRansom(t *testing.T) {
	testcases := []TestCase{
		{String1: "give me one grand today night", String2: "give one grand today", Output: "Yes"},
		{String1: "two times three is not four", String2: "two times two is four", Output: "No"},
		{String1: "ive got a lovely bunch of coconuts", String2: "ive got some coconuts", Output: "No"},
		{String1: "apgo clm w lxkvg mwz elo bg elo lxkvg elo apgo apgo w elo bg", String2: "elo lxkvg bg mwz clm w", Output: "Yes"},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("Test Case : %d\n", (i+1)), func(test *testing.T) {
			output := Ransom(testcase.String1, testcase.String2)

			if !reflect.DeepEqual(output, testcase.Output) {
				test.Errorf("expected %+v, got %+v", testcase.Output, output)
			}
		})
	}
}
