package algolp

import "fmt"

type DoubleLinkedList struct {
	v int
	next *DoubleLinkedList
	prev *DoubleLinkedList
}

func (l DoubleLinkedList) Head() int {
	return l.v
}

func (l DoubleLinkedList) Tail() DoubleLinkedList {
	return *(l.next)
}

func GetEmptyDoubleList() DoubleLinkedList{
	return DoubleLinkedList{next: nil, prev: nil}
}

func AppendDouble(val int, l DoubleLinkedList) DoubleLinkedList{
	var new DoubleLinkedList = DoubleLinkedList{v: val, next: &l, prev: nil}
	l.prev = &new
	return new
}

func (l DoubleLinkedList) Affiche() string{
	var s string
	for !l.IsEmpty() {
		s = fmt.Sprint(s, " ", l.Head())
		l = l.Tail()
	}
	return s
}

func (l DoubleLinkedList) IsEmpty() bool{
	return l.next == nil
}