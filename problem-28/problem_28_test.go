package problem_28

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCases struct {
	MainString string
	Substring  string
	Expected   int
}

func TestProblem28(t *testing.T) {
	testCases := []TestCases{
		{MainString: "leetcode", Substring: "leet", Expected: 0},
		{MainString: "leetcode", Substring: "leeo", Expected: -1},
		{MainString: "leetcode", Substring: "etc", Expected: 2},
		{MainString: "abc", Substring: "c", Expected: 2},
		{MainString: "aaa", Substring: "aaa", Expected: 0},
		{MainString: "mississippi", Substring: "issip", Expected: 4},
		{MainString: "leeteetcode", Substring: "eetc", Expected: 4},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase: %d\n", i+1), func(test *testing.T) {
			output := Problem28V2(testCase.MainString, testCase.Substring)
			if !reflect.DeepEqual(output, testCase.Expected) {
				test.Errorf("Expected : %d Got : %d\n", testCase.Expected, output)
			}
		})
	}
}
