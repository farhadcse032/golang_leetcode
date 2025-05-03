package bubblesort

import (
	"fmt"
	"reflect"
	"testing"
)


type TestCase struct{
	InputArray []int32
	ExpectedArray []int32
}


func TestBubblesort(t *testing.T){

	testCases:=[]TestCase{
		{
			InputArray: []int32{8,7,9,1,3,2,5,10,4,20},
			ExpectedArray: []int32{1,2,3,4,5,7,8,9,10,20},
		},
	}

	for index, testCase := range testCases{

		t.Run(fmt.Sprintf("TestCase%d:\n",index+1),func(test *testing.T) {
			output:=Bubblesort(testCase.InputArray)
			if !reflect.DeepEqual(output,testCase.ExpectedArray){
				test.Errorf("Expected : %v , Got : %v", testCase.ExpectedArray,output)
			}
		})
	}
}