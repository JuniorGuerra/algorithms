package test

import (
	"hashmap/solution"
	"reflect"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: testing.CoverMode(), args: args{
			list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		}, want: []string{
			"Shogun",
		}},
		{name: testing.CoverMode(), args: args{
			list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2: []string{"KFC", "Shogun", "Burger King"},
		}, want: []string{
			"Shogun",
		}},
		{name: testing.CoverMode(), args: args{
			list1: []string{"happy", "sad", "good"},
			list2: []string{"sad", "happy", "good"},
		}, want: []string{
			"sad", "happy",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.FindRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
