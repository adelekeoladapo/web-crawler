package queue

import "errors"

type Queue interface {
	Push(item interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Size() int
	Empty() bool
}

type SimpleQueue struct {
	items []interface{}
}

func (o *SimpleQueue) Push(item interface{}) (err error) {
	o.items = append(o.items, item)
	return
}

func (o *SimpleQueue) Pop() (item interface{}, err error) {
	item, err = o.Peek()
	if err != nil {
		return
	}
	o.items = o.items[:len(o.items)-1]
	return
}

func (o *SimpleQueue) Peek() (item interface{}, err error) {
	if o.Size() == 0 {
		err = errors.New("empty que")
		return
	}
	item = o.items[len(o.items)-1]
	return
}

func (o *SimpleQueue) Size() (size int) {
	size = len(o.items)
	return
}

func (o *SimpleQueue) Empty() bool {
	return !(len(o.items) > 0)
}

func GetSimpleQueue() Queue {
	return &SimpleQueue{}
}
