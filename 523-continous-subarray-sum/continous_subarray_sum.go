package continoussubarraysum

func CheckSubarraySum(nums []int, k int) bool {

	remaindermap := map[int]int{0:-1}
	sum := 0

	for i, num := range nums {

		sum += num
		rem := sum

		if k > 0 {
			rem = sum % k
		}

		if previousIndex, exists := remaindermap[rem]; exists {

			if i-previousIndex >= 2 {
				return true
			}
		} else {
			remaindermap[rem] = i
		}
	}

	return false
}
