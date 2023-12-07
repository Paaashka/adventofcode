package main

import "testing"

func Test_fetchNumberFromString(t *testing.T) {
	tests := []struct {
		in   string
		want int64
	}{
		{in: "grdtdczfm5krxslvfk", want: 55},
		{in: "one99xgk2", want: 12},
		{in: "klklfnxcnmczrjlprktwo55", want: 25},
		{in: "ninetwo52", want: 92},
		{in: "two1nine", want: 29},
		{in: "eightwothree", want: 83},
		{in: "abcone2threexyz", want: 13},
		{in: "xtwone3four", want: 24},
		{in: "4nineeightseven2", want: 42},
		{in: "zoneight234", want: 14},
		{in: "7pqrstsixteen", want: 76},
		{in: "tdjkrtrdj7twoneg", want: 71},
		{in: "9twonexr", want: 91},
		{in: "9bpzdrrfqcs7eightwob", want: 92},
		{in: "jnccdbplkfq6oneightd", want: 68},
		{in: "38oneightg", want: 38},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := fetchNumberFromString(tt.in); got != tt.want {
				t.Errorf("fetchNumberFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
