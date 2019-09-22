package main

import "sync"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	head *Node
	end  *Node
	lock sync.RWMutex
	count int
}

func NewQueue() *Queue {
	q := &Queue{head: nil, end :nil}
	return q
}

func (q *Queue) push(data interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	n := &Node{data: data, next: nil}

	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.count += 1
	return
}

func (q *Queue) pop() (interface{}, bool) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.head == nil {
		return nil, false
	}

	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
	q.count -= 1
	return data, true
}

func (q *Queue) peek() (interface{}, bool) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.head == nil {
		return nil, false
	}

	data := q.head.data
	return data, true
}

func (q *Queue) GetCount()int {
	return q.count
}
