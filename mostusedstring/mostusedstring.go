package mostusedstring

import (
	"fmt"
	"strings"
)


func CheckMostUsed(ss string) {
	m := make(map[int]string)
	j := 0

	var temp strings.Builder
	for i := 0; i < len(ss); i++ {

		if ss[i] == ' ' || ss[i] == '.' {
			if len(temp.String()) > 0 {
				m[j] = temp.String()
				j++
			}
			temp.Reset()
			continue
		}

		temp.WriteByte(ss[i])
	}
	fmt.Printf("%v", m)
	finalmap := make(map[string]int)

	for k := 0; k < len(m); k++ {
		for l :=k; l < len(m); l++ {
			if m[k] == m[l] && m[k] != string(' ') {
				sss := m[k]
				finalmap[sss] += 1

				fmt.Println(sss,finalmap[sss],"|",l,k)
			}
		}
	}

	fmt.Println("output:", finalmap)
}