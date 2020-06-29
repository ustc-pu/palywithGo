package queue

type Queue []interface{}

func (q *Queue) Push(val interface{}) {
	*q = append(*q, val)
}

func (q *Queue) Pop() interface{} {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

//func (q *Queue) Pop() int {
//	res := (*q)[0]
//	*q = (*q)[1:]
//	return res.(int)
//}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}


