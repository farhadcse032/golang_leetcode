package removeduplicatefromsortedarray

import "fmt"




//     A new slice allocation every time.
//     A copy of elements after index i+1 to index i.
//     Reducing the slice length, then adjusting the loop index.
// Each duplicate removal is O(n) due to copying — and if you do that inside a loop, it becomes O(n²) in worst case.
func RemoveDuplicatesv1(nums []int) int {
    
    for i:=0;i <len(nums)-1;i++{

        if (nums[i]==nums[i+1]){
            nums=append(nums[:i],nums[:i+1]...)
            i--
        }
    }
    
    return len(nums)
}

//Time complexity optimized solution
func RemoveDuplicate(nums []int) int {

	temp := nums[0]
	j := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != temp {
			nums[j] = nums[i]
			temp = nums[i]
			j++
		}
		
	}
	fmt.Println(nums)
	return len(nums)

}
