package leetcode

import (
	"strings"

	"github.com/emirpasic/gods/v2/maps/hashmap"
	"github.com/samber/lo"
)

func findKeyPair(key string, stack *hashmap.Map[string, string]) (string, string, bool) {
	k, v, f := "", "", false
	kList := stack.Keys()

	if len(kList) == 0 {
		return k, v, f
	}

	value, found := stack.Get(key)
	if found {
		return key, value, true
	}

	return k, v, f
}

func wordPattern(pattern string, s string) bool {
	split_pattern := strings.Split(pattern, "")
	split_s := strings.Split(s, " ")
	result := true

	if len(split_pattern) != len(split_s) {
		return false
	}

	stack := hashmap.New[string, string]()

	for index := range len(split_pattern) {
		key := split_pattern[index]
		value := split_s[index]
		stack_key, stack_value, found := findKeyPair(key, stack)

		if found {
			if key == stack_key && value == stack_value {
				continue
			}
			if key != stack_key && value == stack_value || key == stack_key && value != stack_value {
				return false
			}
		} else {
			stack.Put(key, value)
		}
	}

	dup := lo.Uniq(split_s)

	if stack.Size() != len(dup) {
		result = false
	}

	return result
}
