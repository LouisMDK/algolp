package algolp

import (
	"testing"
)

func TestPileVide(t *testing.T) {
	var p Pile = GetEmptyPile()
	var f Pile = GetEmptyPile()
	f.cont = 2
	if !p.IsEmpty() || f.IsEmpty() {
		t.Fail()
	}
}

func TestPushPile(t *testing.T) {
	var p Pile = GetEmptyPile()
	p.Push(5)
	if p.element.v != 5 {
		t.Fail()
	}
	p.Push(15)
	if p.element.v != 15 || p.element.next.v != 5 {
		t.Fail()
	}	
}

func TestPopPile(t *testing.T) {
	var p Pile = GetEmptyPile()
	p.Push(5)
	p.Push(3)

	var x int
	x = p.Pop()
	if x != 3 || p.element.v != 5{
		t.Fail()
	}
	p.Push(12)
	p.Push(7)

	x = p.Pop()
	if x != 7 || p.element.v != 12 || p.element.next.v != 5 || p.Count() != 2 {
		t.Fail()
	}
}
