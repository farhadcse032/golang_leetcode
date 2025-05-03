// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	res := sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	fmt.Println(res)
}

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here

	mmap := make(map[int32]int32)
	var count int32

	for _, v := range ar {
		mmap[v]++
		if mmap[v]%2==0{
			count++
		}
	}
	// for _, v := range mmap {
	// 	count +=v / 2
	// }
	return count
}
