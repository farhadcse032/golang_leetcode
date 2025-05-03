package problem34

//{GivenArray:[]int{5,7,7,8,8,10},Target:8,Expected:[]int{3,4},},

func SearchRange(nums []int, target int) []int {
	f,s:=-1,-1
	i:=0
	j:=len(nums)-1
	for j>=0{
		if(nums[i]==target && f==-1){
			f=i
		}
		if(nums[j]==target && s==-1){
			
			s=j
		}
		if (f>-1 && s>-1){
			return []int{f,s}
		}
		i++
		j--
	}
	return []int{f, s}
}
