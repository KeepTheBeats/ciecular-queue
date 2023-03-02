package circularqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircularQueue(t *testing.T) {
	_, err := NewCircularQueue(0)
	assert.Error(t, err)
}

func TestIdxPlusOne(t *testing.T) {
	cq, err := NewCircularQueue(3)
	assert.NoError(t, err)

	assert.Equal(t, 1, cq.IdxPlusOne(0))
	assert.Equal(t, 2, cq.IdxPlusOne(1))
	assert.Equal(t, 0, cq.IdxPlusOne(2))

	cq, err = NewCircularQueue(5)
	assert.NoError(t, err)

	assert.Equal(t, 1, cq.IdxPlusOne(0))
	assert.Equal(t, 2, cq.IdxPlusOne(1))
	assert.Equal(t, 0, cq.IdxPlusOne(4))
}

func TestIsEmptyEnqueueIsFull(t *testing.T) {
	// 0 item
	cq, err := NewCircularQueue(6)
	assert.NoError(t, err)
	assert.Equal(t, true, cq.IsEmpty())
	assert.Equal(t, false, cq.IsFull())

	// 1 - 4 items
	for i := 0; i < cq.Capacity-2; i++ {
		err = cq.Enqueue(i)
		t.Log(cq.Rear, cq.Front)
		assert.NoError(t, err)
		assert.Equal(t, false, cq.IsEmpty())
		assert.Equal(t, false, cq.IsFull())
	}

	// 5 items, full
	err = cq.Enqueue("test enqueue string")
	assert.NoError(t, err)
	assert.Equal(t, false, cq.IsEmpty())
	assert.Equal(t, true, cq.IsFull())

	// full, cannot enqueue
	err = cq.Enqueue(1)
	assert.Error(t, err)
	assert.Equal(t, false, cq.IsEmpty())
	assert.Equal(t, true, cq.IsFull())

	t.Log(cq.Queue)
}

func TestEnqueueDequeueShowItems(t *testing.T) {
	cq, _ := NewCircularQueue(5)
	t.Log(cq.ShowItems())

	cq.Enqueue(10)
	t.Log(cq.ShowItems())

	cq.Enqueue(2.5)
	t.Log(cq.ShowItems())

	cq.Enqueue("aa")
	t.Log(cq.ShowItems())

	cq.Enqueue("bb")
	t.Log(cq.ShowItems())

	item, _ := cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	cq.Enqueue('c')
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	cq.Enqueue(false)
	t.Log(cq.ShowItems())

	cq.Enqueue(5)
	t.Log(cq.ShowItems())

	cq.Enqueue(5)
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	item, _ = cq.Dequeue()
	t.Log("dequeue: ", item)
	t.Log(cq.ShowItems())

	cq.Enqueue(5)
	t.Log(cq.ShowItems())

	cq.Enqueue(true)
	t.Log(cq.ShowItems())

	cq.Enqueue(5)
	t.Log(cq.ShowItems())

	cq.Enqueue(true)
	t.Log(cq.ShowItems())

	cq.Enqueue(true)
	t.Log(cq.ShowItems())

	cq.Enqueue(5)
	t.Log(cq.ShowItems())

	cq.Enqueue(5)
	t.Log(cq.ShowItems())
}

func TestEnqueueDequeueIsEmptyIsFull(t *testing.T) {
	n := 1000000

	chItem := make(chan interface{}, 0)
	go func() {
		for i := 0; i < n; i++ {
			select {
			case chItem <- 1:
			case chItem <- 2:
			case chItem <- 1.5:
			case chItem <- 2.5:
			case chItem <- "str1":
			case chItem <- "str2":
			case chItem <- true:
			case chItem <- false:
			}
		}
		close(chItem)
	}()

	chEnorDe := make(chan bool, 0)
	go func() {
		for i := 0; i < n; i++ {
			select {
			case chEnorDe <- true:
			case chEnorDe <- false:
			}
		}
		close(chEnorDe)
	}()

	var items map[interface{}]int = make(map[interface{}]int)
	var enorDe map[bool]int = make(map[bool]int)
	cq, err := NewCircularQueue(5)
	assert.NoError(t, err)
	var itemCount int = 0

	for {
		item, okItem := <-chItem
		en, okEnorDe := <-chEnorDe
		if !okItem && !okEnorDe {
			break
		}

		if okItem && okEnorDe {
			items[item]++
			enorDe[en]++
			if en {
				err := cq.Enqueue(item)
				if itemCount == cq.Capacity-1 {
					//t.Log(err.Error())
					assert.Error(t, err)
				} else {
					itemCount++
					assert.NoError(t, err)
				}
			} else {
				_, err := cq.Dequeue()
				if itemCount == 0 {
					//t.Log(err.Error())
					assert.Error(t, err)
				} else {
					itemCount--
					assert.NoError(t, err)
				}
			}
			continue
		}
		t.Fatalf("okItem is %t, okEnorDe is %t. They are not the same.", okItem, okEnorDe)
	}

	// check the possibility of every value is average
	for item, count := range items {
		t.Log(item, count, float64(count)/1000000.0)
	}
	for en, count := range enorDe {
		t.Log(en, count, float64(count)/1000000.0)
	}
}
