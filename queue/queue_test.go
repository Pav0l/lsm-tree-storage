package queue

import (
	"testing"
)

func assertEql(expected, received interface{}, t *testing.T) {
	if expected != received {
		t.Error("Invalid result, expected:", expected, "received:", received)
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := new(Queue)

	length := q.getLength()
	assertEql(0, int(length), t)

	q.Enqueue(Node{Data: Data{Key: "first", Value: "fu"}})
	length = q.getLength()
	assertEql(1, int(length), t)
	assertEql("first", q.First.Data.Key, t)

	q.Enqueue(Node{Data: Data{Key: "second", Value: "fu"}})
	q.Enqueue(Node{Data: Data{Key: "third", Value: "fu"}})
	q.Enqueue(Node{Data: Data{Key: "forth", Value: "fu"}})
	q.Enqueue(Node{Data: Data{Key: "fifth", Value: "fu"}})
	length = q.getLength()
	assertEql(5, int(length), t)

	// position of first node in queue is unchanged
	assertEql("first", q.First.Data.Key, t)

	assertEql("fifth", q.Last.Data.Key, t)
}

func TestQueue_Dequeue(t *testing.T) {
	q := new(Queue)

	t.Run("calling Dequeue on empty Queue won't panic", func(t *testing.T) {
		key, val := q.Dequeue()
		assertEql("", key, t)
		assertEql("", val, t)
	})

	q.Enqueue(Node{Data: Data{Key: "1", Value: "fu1"}})
	q.Enqueue(Node{Data: Data{Key: "2", Value: "fu2"}})

	t.Run("Dequeue will remove proper node from que", func(t *testing.T) {
		key, val := q.Dequeue()
		assertEql(1, int(q.getLength()), t)
		assertEql("1", key, t)
		assertEql("fu1", val, t)
		assertEql("2", q.First.Data.Key, t)
		assertEql("fu2", q.First.Data.Value, t)
	})

	t.Run("Dequeue on last node will leave First and Last Queue attrs nil", func(t *testing.T) {
		key, val := q.Dequeue()
		assertEql("2", key, t)
		assertEql("fu2", val, t)
		assertEql(0, int(q.getLength()), t)
		if q.First != nil || q.Last != nil {
			t.Error("First and Last nodes should be nil when emptying Queue. Received First:", q.First, "Last:", q.Last)
		}
	})
}
