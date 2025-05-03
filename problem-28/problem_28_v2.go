package problem_28


/*
mainstr="mississippi"
substr="issip"
*/

func Problem28V2(haystack ,needle string) int {

	for index,_:= range haystack{
		//fmt.Println(haystack[index:len(needle)+index])
		if (index>len(haystack)-len(needle)){
			return -1
		}else if (haystack[index:len(needle)+index])==needle {
			return index
		}
	}
	
	return -1
}