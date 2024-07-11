package test

import (
	"hashmap/solution"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: testing.CoverMode(),
			args: args{nums: []int{4, 1, 2, 1, 3, 3, 2}},
			want: 4,
		},

		{
			name: testing.CoverMode(),
			args: args{nums: []int{2, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.SingleNumber(tt.args.nums); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
