package problem_28

import (
	"log"
)

/*
mainstr="mississippi"
substr="issip"
*/

func Problem28(haystack ,needle string) int {
	result:=0
	flag:=0
	for index,_:= range haystack{
		log.Println("index",index,string(haystack[index]))
		if flag==len(needle) || (haystack==needle){
			return result
		}else if string(haystack[index]) == string(needle[flag])  {
			
			if len(needle)==1{
				//log.Println("here1")
				return index
			}else if (flag==0){
				result=index
			}else if flag==len(needle){
				//log.Println("here2")
				return result
			}
			flag+=1
		}else{
			flag=0
		}
	}
	
	return -1
}