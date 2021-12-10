package data

// Dumbest queue impl possible

type Queue struct {
	data []interface{}
}

func (q *Queue) Push(e interface{}) {
	q.data = append(q.data, e)
}

func (q *Queue) Pop() interface{} {
	out := q.data[0]
	q.data = q.data[1:]
	return out
}

func (q *Queue) Len() int {
	return len(q.data)
}
