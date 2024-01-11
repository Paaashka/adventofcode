package main

import "testing"

func Test_applyMaps(t *testing.T) {
	type args struct {
		in int64
		ms []Map
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1",
			args: args{
				in: 81,
				ms: []Map{
					{dest: 88, source: 18, len: 7},
				},
			},
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyMaps(tt.args.in, tt.args.ms); got != tt.want {
				t.Errorf("applyMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
