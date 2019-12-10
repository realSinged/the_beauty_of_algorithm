package queue

//CircularQueue 循环队列
type CircularQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (c *CircularQueue) IsEmpty() bool {
	if c.head == c.tail {
		return true
	}
	return false
}

func (c *CircularQueue) IsFull() bool {
	if c.head == ((c.tail)+1)%c.capacity {
		return true
	}
	return false
}

func (c *CircularQueue) EnQueue(v interface{}) bool {
	if c.IsFull() {
		return false
	}
	c.q[c.tail] = v
	c.tail = (c.tail + 1) % c.capacity
	return true
}

func (c *CircularQueue) Dequeue() interface{} {
	if c.IsEmpty() {
		return nil
	}
	v := c.q[c.head]
	c.head = (c.head + 1) % c.capacity
	return v
}
