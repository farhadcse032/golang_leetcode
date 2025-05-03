package problem347

//		{Nums: []int{1, 1, 1, 2, 2, 3}, K: 2, Output: []int{1, 2}},

func topKFrequent(nums []int, k int) []int {

	if len(nums) == k {
		return []int{1}
	}
	result := make([]int, 0)
	myMap := make(map[int]int)
	for _, value := range nums {
		if _, ok := myMap[value]; ok {
			myMap[value]++
		} else {
			myMap[value] = 1
		}
	}
	return result
}
