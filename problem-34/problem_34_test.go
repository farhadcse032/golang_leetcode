package problem34

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCases struct {
	GivenArray []int
	Target     int
	Expected   []int
}


func TestSearchRange (t *testing.T) {

	testCases:=[]TestCases{
		{GivenArray:[]int{5,7,7,8,8,10},Target:8,Expected:[]int{3,4},},
		{GivenArray:[]int{5,7,7,8,8,8,10},Target:8,Expected:[]int{3,5},},
		{GivenArray:[]int{5,7,7,8,8,10},Target:6,Expected:[]int{-1,-1},},
		{GivenArray:[]int{},Target:0,Expected:[]int{-1,-1},},
		{GivenArray:[]int{1},Target:1,Expected:[]int{0,0},},
	}
	for i,testCase := range testCases{
		t.Run(fmt.Sprintf("Test Case:%d\n",(i+1)),func(test *testing.T) {
			output:=SearchRangeWithLogN(testCase.GivenArray,testCase.Target)
			if !reflect.DeepEqual(output,testCase.Expected) {
					test.Errorf("Test Case %d Failed, expected: %d --> output: %d\n",(i+1), testCase.Expected,output)
			}
		})

	}
}