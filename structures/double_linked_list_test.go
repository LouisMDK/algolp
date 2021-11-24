package algolp

import (
	"testing"
)
func TestDoubleVide(t *testing.T) {
	var l DoubleLinkedList = GetEmptyDoubleList()
	if !l.IsEmpty() {
		t.Fail()
	}
}

func TestDoubleListe(t *testing.T) {
	var l DoubleLinkedList = GetEmptyDoubleList()
	var d DoubleLinkedList = DoubleLinkedList{v: 2, next: &l, prev: nil}

	if d.Tail() != l || d.Head() != 2 || d.prev != nil{
		t.Fail()
	}
}

func TestDoubleListe2(t *testing.T) {
	var l DoubleLinkedList = GetEmptyDoubleList()
	var d DoubleLinkedList = DoubleLinkedList{v: 1, next: &l, prev: nil}
	l.prev = &d
	var f DoubleLinkedList = DoubleLinkedList{v: 2, next: &d, prev: nil}
	d.prev = &f
	if f.Head() != 2 || f.prev != nil || f.next != &d || f.next.Head() != 1 || !f.next.next.IsEmpty(){
		t.Fail()
	}
}

func TestDoubleListAppend(t *testing.T) {
	var e DoubleLinkedList = GetEmptyDoubleList()
	var f DoubleLinkedList = AppendDouble(2, e)
	var g DoubleLinkedList = AppendDouble(3, f)
	var h DoubleLinkedList = AppendDouble(4, g)

	if h.Head() != 4 || h.Tail().Head() != 3 || h.prev != nil{
		t.Fail()
	}
	if h.next.Head() != 3 || h.next.prev.Head() != 4 {
		t.Fail()
	}
	if h.next.next.Head() != 2 || h.next.next.prev.Head() != 3 {
		t.Fail()
	}
	if !h.next.next.next.IsEmpty() || h.next.next.next.prev.Head() != 2{
		t.Fail()
	}
}
