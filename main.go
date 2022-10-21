package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "20000000000000000000"
	// result := myAtoi(s)
	// fmt.Println(result)

	// mostusedstring.CheckMostUsed(ss)

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
