package ds

import (
	"testing"
)

func TestIntQueue_Pop(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "pop 1 item",
			fields: fields{data: []int{1}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intQ := &IntQueue{
				data: tt.fields.data,
			}
			if got := intQ.Pop(); got != tt.want {
				t.Errorf("IntQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		q := NewIntQueue(10)

		for i := 0; i < 100; i++ {
			q.Push(i)
		}

		for i := 0; i < 100; i++ {
			sum += q.Pop()
		}

		_ = sum
	}

}

func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		q := newQueue(10)

		for i := 0; i < 100; i++ {
			q.push(i)
		}

		for i := 0; i < 100; i++ {
			sum += q.pop().(int)
		}

		_ = sum
	}
}
