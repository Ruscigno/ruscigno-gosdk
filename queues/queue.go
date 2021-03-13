// Copyright 2021 Sander Ruscigno
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package queues

import (
	"errors"
	"sync"
)

//Node storage of queue data
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

//QueueBackend Backend storage of the queue, a double linked list
type QueueBackend struct {
	//Pointers to root and end
	head *Node
	tail *Node

	//keep track of current size
	size    uint32
	maxSize uint32
}

func (queue *QueueBackend) createNode(data interface{}) *Node {
	node := Node{}
	node.data = data
	node.next = nil
	node.prev = nil

	return &node
}

func (queue *QueueBackend) put(data interface{}) error {
	if queue.size >= queue.maxSize {
		err := errors.New("queue is full")
		return err
	}

	if queue.size == 0 {
		//new root node
		queue.head = queue.createNode(data)
		queue.tail = queue.head
		queue.size++
		return nil
	}

	//queue non-empty append to head
	newNode := queue.createNode(data)
	oldTail := queue.tail
	newNode.prev = oldTail
	oldTail.next = newNode

	queue.tail = newNode
	queue.size++
	return nil
}

func (queue *QueueBackend) pop() (interface{}, error) {
	if queue.size == 0 {
		err := errors.New("queue is empty")
		return "", err
	}

	currentHead := queue.head
	newHead := currentHead.next
	queue.head = newHead

	if newHead != nil {
		newHead.prev = nil
	}

	queue.size--
	if queue.size == 0 {
		queue.head = nil
		queue.tail = nil
	}

	return currentHead.data, nil
}

func (queue *QueueBackend) isEmpty() bool {
	return queue.size == 0
}

func (queue *QueueBackend) isFull() bool {
	return queue.size >= queue.maxSize
}

//ConcurrentQueue concurrent queue
type ConcurrentQueue struct {
	//mutex lock
	lock *sync.Mutex

	//empty and full locks
	notEmpty *sync.Cond
	notFull  *sync.Cond

	//queue storage backend
	backend *QueueBackend
}
