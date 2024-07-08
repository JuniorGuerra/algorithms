package test

import (
	"hashmap/solution"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: testing.CoverMode(), args: args{
			[]int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			},
		}, want: false},
		{name: testing.CoverMode(), args: args{
			[]int{
				1, 1, 3, 4, 5, 6, 7, 8, 9, 10,
			},
		}, want: true},
		{name: testing.CoverMode(), args: args{
			[]int{
				1, 2, 3, 4, 5, 6, 2, 8, 9, 10,
			},
		}, want: true},
		{name: testing.CoverMode(), args: args{
			[]int{
				1, 2, 3, 3, 5, 6, 2, 8, 9, 10,
			},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.FindDuplicates(tt.args.arr); got != tt.want {
				t.Errorf("FindDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
