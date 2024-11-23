package leetcode

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func countBits(n int) []int {
	v := []int{}

	for num := range n + 1 {
		binary := fmt.Sprintf("%b", num)
		splitBinary := strings.Split(binary, "")
		filter := lo.Filter(splitBinary, func(chr string, index int) bool {
			return chr != "0"
		})
		v = append(v, len(filter))
	}

	return v
}
