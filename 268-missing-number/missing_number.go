package missingnumber

func MissingNumber(nums []int) int {

	n := len(nums)

	nsum := (n * (n + 1)) / 2

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return nsum - sum
}

func missingNumber(nums []int) int {
	n := len(nums)
	res := n

	for i, v := range nums {
		res ^= i ^ v
	}

	return res
}
