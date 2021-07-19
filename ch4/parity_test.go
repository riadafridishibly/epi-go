package ch4

import (
	"testing"
)

func TestParityNaive(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Parity of 7",
			args: args{x: 7},
			want: 1,
		},
		{
			name: "Parity of 6",
			args: args{x: 6},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParityNaive(tt.args.x); got != tt.want {
				t.Errorf("ParityNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParityRemoveLowestSetBit(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Parity of 7",
			args: args{x: 7},
			want: 1,
		},
		{
			name: "Parity of 6",
			args: args{x: 6},
			want: 0,
		},
		{
			name: "Parity of 9",
			args: args{x: 9},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParityRemoveLowestSetBit(tt.args.x); got != tt.want {
				t.Errorf("ParityRemoveLowestSetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParityUsingCache(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Parity of 7",
			args: args{x: 7},
			want: 1,
		},
		{
			name: "Parity of 6",
			args: args{x: 6},
			want: 0,
		},
		{
			name: "Parity of 9",
			args: args{x: 9},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParityUsingCache(tt.args.x); got != tt.want {
				t.Errorf("ParityUsingCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
