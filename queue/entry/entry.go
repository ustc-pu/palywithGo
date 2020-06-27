package main

import (
	"fmt"
	"pcheng/palywithGo/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())
}
