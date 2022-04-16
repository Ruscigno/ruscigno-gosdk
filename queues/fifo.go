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
	"sync"
)

//NewConcurrentQueue Creates a new queue
func NewConcurrentQueue(maxSize uint32) IConcurrentQueue {
	queue := ConcurrentQueue{}

	//init mutexes
	queue.lock = &sync.Mutex{}
	queue.notFull = sync.NewCond(queue.lock)
	queue.notEmpty = sync.NewCond(queue.lock)

	//init backend
	queue.backend = &QueueBackend{}
	queue.backend.size = 0
	queue.backend.head = nil
	queue.backend.tail = nil

	queue.backend.maxSize = maxSize
	return &queue
}

func (c *ConcurrentQueue) Enqueue(data interface{}) error {
	c.lock.Lock()

	for c.backend.isFull() {
		//wait for empty
		c.notFull.Wait()
	}

	//insert
	err := c.backend.put(data)

	//signal notEmpty
	c.notEmpty.Signal()

	c.lock.Unlock()

	return err
}

func (c *ConcurrentQueue) Dequeue() (interface{}, error) {
	c.lock.Lock()

	for c.backend.isEmpty() {
		c.notEmpty.Wait()
	}

	data, err := c.backend.pop()

	//signal notFull
	c.notFull.Signal()

	c.lock.Unlock()

	return data, err
}

func (c *ConcurrentQueue) GetSize() uint32 {
	c.lock.Lock()
	size := c.backend.size
	c.lock.Unlock()

	return size
}
