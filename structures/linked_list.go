package algolp

import "fmt"

type LinkedList struct {
	head int
	tail *LinkedList
}

func (l LinkedList) Head() int {
	return l.head
}

func (l LinkedList) Tail() LinkedList {
	return *(l.tail)
}

func GetEmptyList() LinkedList{
	return LinkedList{tail: nil}
}

func Append(v int, l LinkedList) LinkedList{
	return LinkedList{head: v, tail: &l}
}

func (l LinkedList) Affiche() string{
	var s string
	for !l.IsEmpty() {
		s = fmt.Sprint(s, " ", l.Head())
		l = l.Tail()
	}
	return s
}

func (l LinkedList) IsEmpty() bool{
	return l.tail == nil
}