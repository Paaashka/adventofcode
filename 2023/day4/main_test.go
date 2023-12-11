package main

import (
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	tests := []struct {
		in   string
		want []int
	}{
		{
			in:   "12 34 56",
			want: []int{12, 34, 56},
		},
		{
			in:   "12  4 56",
			want: []int{12, 4, 56},
		},
		{
			in:   " 1 21 53 59 44",
			want: []int{1, 21, 53, 59, 44},
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := split(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}
