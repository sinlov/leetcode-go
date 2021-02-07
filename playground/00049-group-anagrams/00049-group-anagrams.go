package _00049_group_anagrams

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	record := map[string][]string{}
	for _, str := range strs {
		sr := []rune(str)
		sort.Sort(sortRunes(sr))
		tmp := record[string(sr)]
		tmp = append(tmp, str)
		record[string(sr)] = tmp
	}

	for _, v := range record {
		res = append(res, v)
	}

	return res
}
