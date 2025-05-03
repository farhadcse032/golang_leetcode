// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slices"
)

func countingSort(arr []int32) []int32 {
	max := slices.Max(arr)+1
	countArray := make([]int32, max)	
fmt.Println(max)
	//arr := []int32{1, 1, 3, 2, 1}
	for _,v:= range arr{
		
		countArray[v]++
		// fmt.Println(countArray)
	}
	

	res:=make([]int32,0)

	for i,v := range countArray{
		for v>0{
			res = append(res, int32(i))
			v--
		}
	}
	// Write your code here
	return countArray

}


func main() {

	// arr := []int32{1, 1, 3, 2, 1}
	arr := []int32{63,54,17,78,43,70,32,97,16,94,74,18,60,61,35,83,13,56,75,52,70,12,24,37,17,0,16,64,34,81,82,24,69,2,30,61,83,37,97,16,70,53,0,61,12,17,97,67,33,30,49,70,11,40,67,94,84,60,35,58,19,81,16,14,68,46,42,81,75,87,13,84,33,34,14,96,7,59,17,98,79,47,71,75,8,27,73,66,64,12,29,35,80,78,80,6,5,24,49,82}
	res := countingSort(arr)
	fmt.Println("res")
	fmt.Println(res)
}
