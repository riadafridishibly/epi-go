package ch5

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicate(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		wantArr []int
		want    int
	}{
		{
			name:    "Test remove duplicate",
			args:    args{[]int{1, 2, 2, 3, 3, 4, 4, 5}},
			wantArr: []int{1, 2, 3, 4, 5, 4, 4, 5},
			want:    5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteDuplicate(tt.args.a); got != tt.want {
				t.Errorf("DeleteDuplicate() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.wantArr, tt.args.a) {
				t.Errorf("DeleteDuplicate() = %v, want %v", tt.args.a, tt.wantArr)
			}
		})
	}
}
