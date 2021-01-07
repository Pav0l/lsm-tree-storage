package queue

// Queue struct
type Queue struct {
	Length uint
	First  *Node
	Last   *Node
}

// Node holds the data and pointer to next node in queue
type Node struct {
	Data Data
	next *Node
}

// Data holds the basic key and value strings
type Data struct {
	Key, Value string
}

// Enqueue - Adds Node to the back of the Queue
func (q *Queue) Enqueue(n Node) {
	if q.Last != nil {
		q.Last.next = &n
	}

	q.addLast(&n)

	if q.First == nil {
		q.addFirst(q.Last)
	}

	q.Length++
}

// Dequeue - Removes the first Node from the Queue and returns its key & value
func (q *Queue) Dequeue() (key, value string) {
	if q.First == nil {
		return "", ""
	}
	first := q.First
	data := first.Data
	q.addFirst(first.next)

	if q.First == nil {
		q.Last = nil
	}
	q.Length--
	return data.Key, data.Value
}

func (q *Queue) getLength() uint {
	return q.Length
}

func (q *Queue) addFirst(n *Node) {
	q.First = n
}

func (q *Queue) addLast(n *Node) {
	q.Last = n
}
