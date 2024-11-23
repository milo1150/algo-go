package leetcode

import (
	"strings"

	"github.com/emirpasic/gods/v2/lists/arraylist"
	"github.com/emirpasic/gods/v2/stacks/arraystack"
	"github.com/samber/lo"
)

func reverseVowels(s string) string {
	r := ""
	vowels := []string{"a", "e", "i", "o", "u"}

	stack := arraylist.New[string]()
	for i := range s {
		char := string(s[i])
		if lo.Contains(vowels, strings.ToLower(char)) {
			stack.Add(char)
		}
	}

	cpStack := arraystack.New[string]()
	for i := range stack.Size() {
		v, _ := stack.Get(i)
		cpStack.Push(v)
	}

	for i := range s {
		char := string(s[i])
		if lo.Contains(stack.Values(), char) {
			pop, _ := cpStack.Pop()
			r += pop
		} else {
			r += char
		}
	}

	return r
}
