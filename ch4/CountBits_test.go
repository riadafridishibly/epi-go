package ch4

import (
	"math"
	"testing"
)

func TestCountBits(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test number of bits in 7",
			args: args{x: 7},
			want: 3,
		},
		{
			name: "Test number of bits in 0",
			args: args{x: 0},
			want: 0,
		},
		{
			name: "Test number of bits in 16",
			args: args{x: 16},
			want: 1,
		},
		{
			name: "Test number of bits in int32 max",
			args: args{x: math.MaxInt32},
			want: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBits(tt.args.x); got != tt.want {
				t.Errorf("CountBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
