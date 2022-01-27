package queue

import "errors"

type Queue interface {
	Enqueue(item interface{})
	Dequeue() (interface{}, error)
	Empty() bool
	Size() int
}

type SimpleQueue struct {
	items []interface{}
}

func (o *SimpleQueue) Enqueue(item interface{}) {
	o.items = append(o.items, item)
}

func (o *SimpleQueue) Dequeue() (item interface{}, err error) {
	if o.Empty() {
		err = errors.New("empty queue")
		return
	}
	item = o.items[0]
	o.items = o.items[1:]
	return
}

func (o *SimpleQueue) Empty() bool {
	return len(o.items) == 0
}

func (o *SimpleQueue) Size() int {
	return len(o.items)
}

func GetSimpleQueue() Queue {
	return &SimpleQueue{}
}
