package palindromicsubstrings


func Palindromicsubstrings(s string) int {

	n:=len(s)
	count:=0

	expand:=func (left,right int){
		for (left >=0 && right < n && s[left] ==s[right]){
			count++
			left--
			right++
		}
	}
	
	for i:=0;i<n;i++{
		expand(i,i)
		expand(i,i+1)
	}

	return count
}

