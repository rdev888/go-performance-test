package main

import (
	"fmt"
	"go-performance-test/DataStructures"
	"runtime"
)

func main() {
	queue := DataStructures.Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Peek())
	fmt.Println(runtime.NumCPU())
}
