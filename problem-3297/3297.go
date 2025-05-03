package problem3297

// You can edit this code!
// Click here and start typing.
//3297. Count Substrings That Can Be Rearranged to Contain a String I
/*
word1 := "bcca", word2 := "abc" output:=1
word1 = "abcabc", word2 = "abc" 10
word1 = "abcabc", word2 = "aaabc" 0
*/

func ValidSubstringCount(word1 string, word2 string) int64 {
	if len(word2) > len(word1) {
		return 0
	}

	R := 0
	tmap := make(map[byte]int)

	count := 0

	for index, _ := range word2 {
		tmap[word2[index]]++
	}
	required := len(tmap)

	for R < len(word1) {
		wmap := make(map[byte]int)
		formed := 0
		L := R
		for L < len(word1) {
			wmap[word1[L]]++
			if tmap[word1[L]] > 0 && wmap[word1[L]] == tmap[word1[L]] {
				formed++
			}
			if formed == required {
				count = count + len(word1) - L
				break
			}

			L++

		}

		R++
	}
	return int64(count)
}

func ValidSubstringCount2(word1 string, word2 string) (ans int64) {
	if len(word1) < len(word2) {
		return 0
	}
	cnt := [26]int{}
	need := 0
	for _, c := range word2 {
		cnt[c-'a']++
		if cnt[c-'a'] == 1 {
			need++
		}
	}
	win := [26]int{}
	l := 0
	for _, c := range word1 {
		i := int(c - 'a')
		win[i]++
		if win[i] == cnt[i] {
			need--
		}
		for need == 0 {
			i = int(word1[l] - 'a')
			if win[i] == cnt[i] {
				need++
			}
			win[i]--
			l++
		}
		ans += int64(l)
	}
	return
}