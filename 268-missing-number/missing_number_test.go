package missingnumber

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct{
	Input []int
	Expected int
}

func TestMissingNumber(t *testing.T){

	testCases:=[]TestCase{
		{
			Input: []int{3,0,1},
			Expected:2,
		},
		{
			Input: []int{0,1},
			Expected:2,
		},{
			Input: []int{9,6,4,2,3,5,7,0,1},
			Expected:8,
		},
	}


	for index, testCase := range testCases{

		t.Run(fmt.Sprintf("TestCase: %d\n",index+1),func( t *testing.T){

			output:=MissingNumber(testCase.Input)

			if !reflect.DeepEqual(output,testCase.Expected){
				t.Errorf(" Expected: %d | Got: %d\n",testCase.Expected,output)
			}
		})
	}
}