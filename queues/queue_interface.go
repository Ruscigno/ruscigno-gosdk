package queues

type IConcurrentQueue interface {
	Enqueue(data interface{}) error
	Dequeue() (interface{}, error)
	GetSize() uint32
}
