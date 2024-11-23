package leetcode

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case 1", args: args{pattern: "abba", s: "dog cat cat dog"}, want: true},
		{name: "case 2", args: args{pattern: "abba", s: "dog cat cat fish"}, want: false},
		{name: "case 3", args: args{pattern: "aaaa", s: "dog cat cat dog"}, want: false},
		{name: "case 4", args: args{pattern: "abba", s: "dog dog dog dog"}, want: false},
		{name: "case 5", args: args{pattern: "aba", s: "dog cat cat"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
