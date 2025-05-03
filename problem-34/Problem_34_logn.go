package problem34


//{GivenArray:[]int{5,7,7,8,8,10},Target:8,Expected:[]int{3,4},},
/*
low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
*/

//5,7,7,8,8,8,10

func SearchRangeWithLogN(nums []int, target int) []int {
    n := len(nums)
    i := lowerBound(nums, target)

    if i == n || nums[i] > target {
        return []int{-1, -1}
    }

    j := lowerBound(nums, target + 1)
    return []int{i, j - 1}
}

func lowerBound(nums []int, target int) int {
    low, high := 0, len(nums)

    for low < high {
        middle := low + (high - low) >> 1

        if nums[middle] < target {
            low = middle + 1
        } else {
            high = middle
        }
    }

    return low
}
