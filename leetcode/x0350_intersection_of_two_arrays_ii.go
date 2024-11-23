package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	r := []int{}
	freq := make(map[int]int)

	for _, num := range nums2 {
		freq[num] += 1
	}

	for _, num := range nums1 {
		find := freq[num]
		if find != 0 {
			r = append(r, num)
			freq[num] -= 1
		}
	}

	return r
}
