package maxsubarraysum

// Time complexity O(n^2)
// Space complexity O(1)
func MaxSubArraySum(nums []int) int {

	// nums:=[-2,1,-3,4,-1,2,1,-5,4]
	var output int
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > output {
				output = sum
			}
		}

	}
	return output
}

// Sliding Window
// Time complexity O(n)
// Space complexity O(1)
func MaxSubArraySumV2(nums []int) int {

	// nums:=[-2,1,-3,4,-1,2,1,-5,4]
	R, L, sum := 0, 0, 0
	output := 0
	for R < len(nums) {
		sum += nums[R]
		if sum > output {
			output = sum
		}
		for sum < 0 {

			sum -= nums[L]
			L++
		}

		R++
	}
	return output
}

// Kadanes Algorithm
// Time complexity O(n)
// Space complexity O(1)

func MaxSubArraySumV3(nums []int) int {

	currsum := nums[0]
	maxsum := nums[0]
	for i :=1;i<len(nums);i++{
	  currsum = max(nums[i],currsum + nums[i])
	  maxsum = max(currsum,maxsum)
	} 
	return maxsum
}
