package leetcode

import (
	"github.com/emirpasic/gods/v2/lists/arraylist"
)

func MoveZeroes(nums []int) {
	result := arraylist.New[int]()
	currentIndex := 0

	for i := range len(nums) {
		if nums[i] == 0 {
			result.Add(0)
		} else {
			result.Insert(currentIndex, nums[i])
			currentIndex += 1
		}
	}

	for j := range result.Size() {
		value, _ := result.Get(j)
		nums[j] = value
	}
}
