package maxsumsubarray

// Sliding Window
// Time complexity O(n)
// Space complexity O(1)
func MaxSumSubArrayV1(nums []int) []int {

	// nums:=[-2,1,-3,4,-1,2,1,-5,4]
	R, L, sum := 0, 0, 0

	minL, maxR := 0, 0
	output := 0
	for R < len(nums) {
		sum += nums[R]

		for sum < 0 {

			sum -= nums[L]
			L++
		}
		if sum > output {
			maxR = R
			minL = L
			output = sum
		}
		R++
	}
	return nums[minL : maxR+1]
}

// Kadanes Algorithm
// Time complexity O(n)
// Space complexity O(1)

func MaxSumSubArrayV2(nums []int) []int {

	// nums:=[-2,1,-3,4,-1,2,1,-5,4]
	x := nums[0]
	currentSum, maxSum := x, x
	min,max,temp:=0,0,0
	for i := 1; i < len(nums); i++ {
		// currentSum = math.Max(float64(nums[i]), float64(currentSum+float64(nums[i])))

		// maxSum = math.Max(currentSum, maxSum)

		if nums[i] > currentSum+ nums[i]{
			currentSum=nums[i]
			temp=i
		}else{
			currentSum+= nums[i]
		}

		if currentSum >maxSum{
			maxSum=currentSum
			min=temp
			max=i
		}
	}
	return nums[min:max+1]
}
