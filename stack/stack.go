package stack

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(items ...int) {
	*s = append(*s, items...)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item
}
