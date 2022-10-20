package twosum


/*

Let [2,4,3,7]
Target 9


1. Make a temporary map
2. Iterate over the number array
3. Check tempArray[9-2] == tempArray[7] exists or not on map
4. If exist on map then return iteration difference 7 and index=>value 2 .
5. If not exist then store tempArray[index=>value 2]= difference

=>Target 9

=>> For 2  diff 9-2=7
temmap[7];not exist
temmap[2]=7

=>> For 4  diff 9-4=5
temmap[5];not exist
temmap[4]=5

=>> For 3  diff 9-3=6
temmap[6];not exist
temmap[3]=6

=>> For 7  diff 9-7=2
temmap[2]; exist ???
 ====> return 2,7

*/

func TwoSum(numArray []int, target int) []int {

	temmap := make(map[int]int)

	for _, v := range numArray {
		diff := target - v
		if _, ok := temmap[diff]; ok {

			return []int{diff, v}
		}
		temmap[v] = diff
	}

	return []int{-1, -1}
}

func TwoSumIndex(numArray []int, target int) []int {

	temmap := make(map[int]int)

	for k, v := range numArray {
		//[]int{1, 3, 4, 5}, 7
		diff := target - v
		if index, ok := temmap[diff]; ok {
			// fmt.Printf("%+v==%+v",temmap,index)
			return []int{index, k}
		}

		// index1, ok1 := temmap[diff]
		temmap[v] = k
	}

	return []int{-1, -1}
}
