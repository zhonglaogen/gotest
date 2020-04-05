package queue

type Queue []interface{}

func (q *Queue) Push(v int)  {

	*q = append(*q, v)
	//*q = append(*q, v.(int))

}

func (q *Queue) Pop() int {

	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool  {

	return len(*q) == 0
}