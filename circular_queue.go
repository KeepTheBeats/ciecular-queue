package circularqueue

import "fmt"

type CircularQueue struct {
	Queue    []interface{}
	Capacity int
	Front    int
	Rear     int
}

func NewCircularQueue(capacity int) (*CircularQueue, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("invalid value %d as the capacity of a queue", capacity)
	}
	return &CircularQueue{
		Queue:    make([]interface{}, capacity, capacity),
		Capacity: capacity,
		Front:    0,
		Rear:     0,
	}, nil
}

func (cq *CircularQueue) IdxPlusOne(idx int) int {
	return (idx + 1) % cq.Capacity
}

// According to the wiki: https://en.wikipedia.org/wiki/Circular_buffer,
// There are 2 solutions to tell the difference between the full and empty:
// 1. The buffer only has a maximum in-use size of Length - 1. The buffer is empty if the start and end indexes are equal and full when the in-use size is Length - 1. (We use this solution.)
// 2. There should be another integer count that is incremented at a write operation and decremented at a read operation. Then checking for emptiness means testing count equals 0 and checking for fullness means testing count equals Length.
func (cq *CircularQueue) IsEmpty() bool {
	return cq.Front == cq.Rear
}

func (cq *CircularQueue) IsFull() bool {
	return cq.IdxPlusOne(cq.Rear) == cq.Front
}

func (cq *CircularQueue) Enqueue(item interface{}) error {
	if cq.IsFull() {
		return fmt.Errorf("the circular queue is full, cannot enqueue")
	}
	cq.Queue[cq.Rear] = item
	cq.Rear = cq.IdxPlusOne(cq.Rear)
	return nil
}

func (cq *CircularQueue) Dequeue() (interface{}, error) {
	if cq.IsEmpty() {
		return nil, fmt.Errorf("the circular queue is empty, cannot dequeue")
	}
	outItem := cq.Queue[cq.Front]
	cq.Front = cq.IdxPlusOne(cq.Front)
	return outItem, nil
}

func (cq *CircularQueue) ShowItems() string {
	var output string
	for idx := cq.Front; idx != cq.Rear; idx = cq.IdxPlusOne(idx) {
		output = fmt.Sprintf("%s %v", output, cq.Queue[idx])
	}
	return output
}
