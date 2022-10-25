package ransom

import (
	"strings"
)

func Ransom(String1, String2 string) string {

	newString1 := strings.Split(String1, " ")
	newString2 := strings.Split(String2, " ")
	var mm = make(map[string]int)
	var ss = make(map[string]int)
	for _, v := range newString1 {
		index, ok := mm[v]
		if ok {
			mm[v] = index + 1
		} else {
			mm[v] = 0
		}

	}

	for _, v := range newString2 {
		index, ok := ss[v]
		if ok {
			ss[v] = index + 1
		} else {
			ss[v] = 0
		}

	}

	for _, v := range newString2 {
		index, ok := mm[v]
		index1, _ := ss[v]
		if !ok {
			return "No"
		} else if index1 > index {
			return "No"
		}
		ss[v] = 0
	}

	return "Yes"
}
