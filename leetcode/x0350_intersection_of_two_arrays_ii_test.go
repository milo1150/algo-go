package leetcode

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 1", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{name: "case 2", args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
		{name: "case 3", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2}}, want: []int{2}},
		{name: "case 4", args: args{nums1: []int{2, 1}, nums2: []int{1, 2}}, want: []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
