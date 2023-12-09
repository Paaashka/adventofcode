package main

import "testing"

func Test_isSymbol(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{in: "3", want: false},
		{in: "", want: false},
		{in: ".", want: false},
		{in: "#", want: true},
		{in: "/", want: true},
		{in: "#", want: true},
		{in: "*", want: true},
		{in: "-", want: true},
		{in: "+", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := isSymbol(tt.in); got != tt.want {
				t.Errorf("isSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
