package main

import "sync"

type MinStack struct {
	items []data
	size int
	lock sync.RWMutex
}

type data struct{
	val int
	min int
}

// Constructor /** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		items: make([]data, 0),
	}
}

func (s *MinStack) Push(val int)  {
	s.lock.Lock()
	defer s.lock.Unlock()

	item := data{
		val: val, min: val,
	}

	if len(s.items) > 0 {
		prevmin := s.items[s.size -1].min
		// if prev node min val less than current val, set curuent item min to prev val
		if prevmin < item.min {
			item.min = prevmin
		}
	}
	s.size++
	s.items = append(s.items, item)
}

func (s *MinStack) Pop()  {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = s.items[:s.size - 1]
	s.size--
}

func (s *MinStack) Top() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.items[s.size - 1].val
}

func (s *MinStack) GetMin() int {
	return s.items[s.size - 1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
