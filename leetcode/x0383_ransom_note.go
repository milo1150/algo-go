package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	countRansomNote := make(map[string]int)
	countMagazine := make(map[string]int)

	for _, chr := range ransomNote {
		countRansomNote[string(chr)] += 1
	}

	for _, chr := range magazine {
		countMagazine[string(chr)] += 1
	}

	for key, value := range countRansomNote {
		v, found := countMagazine[key]
		if !found {
			return false
		}
		if found && v < value {
			return false
		}
	}

	return true
}
