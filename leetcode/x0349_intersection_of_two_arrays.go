package leetcode

import (
	"github.com/samber/lo"
)

func intersection(nums1 []int, nums2 []int) []int {
	r := lo.Intersect(nums1, nums2)
	u := lo.Uniq(r)
	return u
}
