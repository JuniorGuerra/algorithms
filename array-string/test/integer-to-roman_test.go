package test

import (
	"array-string/solutions"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "M", args: args{
			num: 3749,
		}, want: "MMMDCCXLIX"},
		{name: "L", args: args{
			num: 58,
		}, want: "LVIII"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solutions.IntToRoman(tt.args.num); got != tt.want {
				t.Errorf("IntToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
