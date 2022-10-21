package problem219

import (
	"math"
)

/********************************
Given an integer array nums and an integer k,
 return true if there are two distinct indices i and j in the array such that
  nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true	"fmt"
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

/*

# Algorithm

1. Get Map
2. Iterate through the array
3. Check Map[value] Exists ?
							IF Exists then check 
									IF current postion - previous postion <= given ====> TRUE return
									ELSE Map[value]= iteration index (i) Current position as previous position
							ELSE 
								Map[value]=i // Store value as key and postiion as value in Map.
4. FALSE return

{Input: []int{1, 2, 3, 1}, Given: 3, Output: true},
{Input: []int{1, 0, 1, 1}, Given: 1, Output: true},


*/
func Problem219(inputs []int, given int) bool {
	myMap := make(map[int]int)

	for i, v := range inputs {

		if index, ok := myMap[v]; ok {
			// fmt.Println(index,i,v)
			if math.Abs(float64(index)-float64(i)) <= float64(given) {
				return true
			}
			myMap[v] = i
		}
		myMap[v] = i
	}

	return false
}
