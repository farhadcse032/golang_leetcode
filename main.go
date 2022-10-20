package main

import (
	"fmt"

	"go_practise/stock121"
	"go_practise/twosum"
	"strings"
)

func main() {
	// s := "20000000000000000000"
	// result := myAtoi(s)
	// fmt.Println(result)

	// ss := "Hi MD Farhad Hossain. My name is Shova. Actually I know you for a long time. You are not so good whatever I think I am in love with you. Do you accept me if you are interested ?"
	// mostusedstring.CheckMostUsed(ss)
	twosum.TwoSumIndex([]int{1, 3, 4, 5}, 7)

	res1 := stock121.Profit([]int{7, 1, 5, 3, 6, 4})
	res2 := stock121.Profit([]int{2,4,1})
	res3 := stock121.Profit([]int{2,4,2,1})
	res4 := stock121.Profit([]int{2,1,2,0,1})

	fmt.Println("Stock121 Input1: ", res1)
	fmt.Println("Stock121 Input2: ", res2)
	fmt.Println("Stock121 Input3: ", res3)
	fmt.Println("Stock121 Input4: ", res4)
}

func myAtoi(s string) int {

	if len(s) < 0 || len(s) > 200 || s == "" {
		return 0
	}
	var filter strings.Builder
	flag := 0
	if ('0' <= s[0] && s[0] <= '9') || (s[0] == '-') || (s[0] == '+') || (s[0] == ' ') {
		flag = 1
	} else {
		flag = 0
	}
	if len(s) > 1 {
		if ('0' <= s[1] && s[1] <= '9') || (s[1] == ' ') || (s[1] == '.') {
			flag = 1
		} else {
			flag = 0
		}
	}

	if flag == 0 {
		return 0
	}
	num := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('0' <= b && b <= '9') || (b == '-') || (b == '+') {
			filter.WriteByte(b)
			num = 1

		} else if (b == ' ') && num == 1 {
			break
		} else if b == ' ' {
			continue
		} else if b == '.' {
			break
		} else {
			break
		}
	}
	res := filter.String()

	var resInt int64

	fmt.Sscanf(res, "%d", &resInt)
	fmt.Println(resInt)
	if resInt < -2147483648 {
		return -2147483648
	}

	if resInt > 2147483647 {
		return 2147483647
	}
	return int(resInt)

}
