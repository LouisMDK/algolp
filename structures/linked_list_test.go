package algolp

import (
	"testing"
)

func TestVide(t *testing.T) {
	var l LinkedList = GetEmptyList()
	if !l.IsEmpty() {
		t.Fail()
	}
}

func TestListe(t *testing.T) {
	var l LinkedList = GetEmptyList()
	var d LinkedList = LinkedList{head: 2, tail: &l}

	if d.Tail() != l || d.Head() != 2 {
		t.Fail()
	}
}

func TestListe2(t *testing.T) {
	var l LinkedList = GetEmptyList()
	var d LinkedList = LinkedList{head: 2, tail: &l}
	var f LinkedList = LinkedList{head: 3, tail: &d}

	if f.Tail() != d || f.Tail().Tail() != l {
		t.Fail()
	}
}

func TestListeAppend(t *testing.T) {
	var e LinkedList = GetEmptyList()

	e = Append(2, e)
	e = Append(3, e)
	e = Append(4, e)

	if e.Head() != 4 || e.Tail().Head() != 3 || e.Tail().Tail().Head() != 2 || !e.Tail().Tail().Tail().IsEmpty() {
		t.Fail()
	}
}

func TestListeAffiche(t *testing.T) {
	var e LinkedList = GetEmptyList()

	e = Append(3, e)
	e = Append(7, e)
	e = Append(-2, e)
	e = Append(12, e)
	if len(e.Affiche()) != 10 {
		t.Fail()
	}
}
