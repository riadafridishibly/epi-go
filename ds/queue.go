package ds

type queue struct {
	data []interface{}
}

func newQueue(n int) *queue {
	return &queue{
		data: make([]interface{}, 0, n),
	}
}

func (q *queue) length() int {
	return len(q.data)
}

func (q *queue) isEmpty() bool {
	return q.length() == 0
}

func (q *queue) push(v interface{}) {
	q.data = append(q.data, v)
}

func (q *queue) pop() interface{} {
	ret := q.data[0]
	q.data = q.data[1:]

	return ret
}

func (q *queue) peek() interface{} {
	ret := q.data[0]

	return ret
}

type IntQueue struct {
	data []int
}

func NewIntQueue(size int) *IntQueue {
	return &IntQueue{
		data: make([]int, 0, size),
	}
}

func (intQ *IntQueue) Len() int {
	return len(intQ.data)
}

func (intQ *IntQueue) IsEmpty() bool {
	return intQ.Len() == 0
}

func (intQ *IntQueue) Push(value int) {
	intQ.data = append(intQ.data, value)
}

func (intQ *IntQueue) Pop() int {
	ret := intQ.data[0]

	intQ.data = intQ.data[1:]

	return ret
}

func (intQ *IntQueue) Peek() int {
	return intQ.data[0]
}
