package stack

import "errors"

type Stack interface {
	Push(item interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Size() int
	Empty() bool
}

type SimpleStack struct {
	items []interface{}
}

func (o *SimpleStack) Push(item interface{}) {
	o.items = append(o.items, item)
	return
}

func (o *SimpleStack) Pop() (item interface{}, err error) {
	item, err = o.Peek()
	if err != nil {
		return
	}
	o.items = o.items[:len(o.items)-1]
	return
}

func (o *SimpleStack) Peek() (item interface{}, err error) {
	if o.Size() == 0 {
		err = errors.New("empty que")
		return
	}
	item = o.items[len(o.items)-1]
	return
}

func (o *SimpleStack) Size() (size int) {
	size = len(o.items)
	return
}

func (o *SimpleStack) Empty() bool {
	return !(len(o.items) > 0)
}

func GetSimpleStack() Stack {
	return &SimpleStack{}
}
