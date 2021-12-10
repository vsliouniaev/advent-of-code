package data

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(e interface{}) {
	s.data = append(s.data, e)
}

func (s *Stack) Pop() interface{} {
	d := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return d
}

func (s *Stack) Len() int {
	return len(s.data)
}
