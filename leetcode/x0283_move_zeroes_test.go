package leetcode

import (
	"reflect"
	"testing"
)

func Test_MoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "moveZeros_1", args: args{nums: []int{0, 1, 0, 3, 12}}, want: []int{1, 3, 12, 0, 0}},
		{name: "moveZeros_1", args: args{nums: []int{0}}, want: []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// run
			MoveZeroes(tt.args.nums)

			// compare modified nums with expected result (tt.want)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeros() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
