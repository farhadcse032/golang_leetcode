package mergesortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct{
	Input1 []int
	M int
	Input2 []int
	N int
	Output []int
}

func TestMergeSortedArray( t *testing.T){

	testCases:=[]TestCase{
		{
			Input1:[]int{1,2,3,0,0,0},
			M:3,
			Input2:[]int{2,5,6},
			N:3,
			Output:[]int{1,2,2,3,5,6},
		},
		{
			Input1:[]int{0},
			M:0,
			Input2:[]int{1},
			N:1,
			Output:[]int{1},
		},
		{
			Input1:[]int{1},
			M:1,
			Input2:[]int{},
			N:0,
			Output:[]int{1},
		},
	}

	for index, v := range testCases{

		t.Run(fmt.Sprintf("Test Case: %d\n",index+1),func (t *testing.T){
			output:= MergeSortedArray(v.Input1,v.M,v.Input2,v.N)

			if !reflect.DeepEqual(output,v.Output){
				t.Errorf("Expected: %d | Got %d \n",v.Output,output)
			}
		})
	}
}