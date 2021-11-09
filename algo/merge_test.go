package algo

import (
	"reflect"
	"testing"
)

func TestMerge2SortedArray(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Simple tests 1",
			args: args{
				a: []int{1, 2, 3},
				b: []int{4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Simple tests 2",
			args: args{
				a: []int{1, 2, 4},
				b: []int{3},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Simple tests 2",
			args: args{
				a: []int{1, 2, 4},
				b: []int{3},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge2SortedArray(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge2SortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge2SortedArrayN(t *testing.T) {
	type args struct {
		a []int
		b []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Sample test 1",
			args: args{
				a: []int{1, 2, 3, 4},
				b: []int{6, 7},
				n: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "Sample test 2",
			args: args{
				a: []int{1, 2, 3, 4},
				b: nil,
				n: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "Sample test 2",
			args: args{
				a: []int{1, 2, 4, 6, 9},
				b: []int{3, 3, 3},
				n: 5,
			},
			want: []int{1, 2, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge2SortedArrayN(tt.args.a, tt.args.b, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge2SortedArrayN() = %v, want %v", got, tt.want)
			}
		})
	}
}
