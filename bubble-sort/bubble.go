package bubblesort

import "fmt"

func Bubblesort(arr []int32) []int32{
	i:=0
	for index,_ := range arr{
		for iindex,_ := range arr{
			i++
			if arr[index] > arr[iindex]{
				arr[index],arr[iindex] = arr[iindex],arr[index]
			}
		}
	}
	fmt.Println(arr)
	return arr
}