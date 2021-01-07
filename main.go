package main

import (
	"fmt"

	"example.com/lsm/memtable"
	"example.com/lsm/queue"
)

func main() {
	// initialize
	writeQue := new(queue.Queue)
	rbt := new(memtable.RedBlackTree)

	fmt.Println("init")

	writeQue.Enqueue(queue.Node{Data: queue.Data{Key: "kung", Value: "fu"}})
	writeQue.Enqueue(queue.Node{Data: queue.Data{Key: "kick", Value: "flip"}})
	writeQue.Enqueue(queue.Node{Data: queue.Data{Key: "feng", Value: "shuei"}})

	var key, value string
	for writeQue.Length > 0 {
		key, value = writeQue.Dequeue()
		fmt.Printf("key: %s, value: %s\n", key, value)

		rbt.Insert(key)
	}

	fmt.Println("rbt size:", rbt.Size)
	fmt.Println("rbt root:", rbt.Root)
	fmt.Println("end")
}
