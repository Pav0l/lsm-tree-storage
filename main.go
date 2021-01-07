package main

import (
	"fmt"

	"example.com/lsm/queue"
)

func main() {
	// initialize
	writeQue := new(queue.Queue)

	fmt.Println("init")

	n := queue.Node{Data: queue.Data{Key: "kung", Value: "fu"}}
	writeQue.Enqueue(n)

	m := queue.Node{Data: queue.Data{Key: "kick", Value: "flip"}}
	writeQue.Enqueue(m)

	var key, value string
	for writeQue.Length > 0 {
		key, value = writeQue.Dequeue()
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
	fmt.Println("end")
}
