package algolp

import (
	"testing"
)

func TestFileVide(t *testing.T) {
	var f File = GetEmptyFile()
	if f.cont != 0 || f.element.prev != nil || f.element.next != nil || f.last == nil {
		t.Fail()
	}
}

func TestFilePush(t *testing.T) {
	var f File = GetEmptyFile()
	
	f.Push(3)
	f.Push(4)
	f.Push(5)
	if f.Count() != 3 || f.element.v != 5 || f.element.next.v != 4 || f.element.next.next.v != 3 || f.last.v != 3 {
		t.Fail()
	}
}

func TestFilePop(t *testing.T) {
	var f File = GetEmptyFile()
	f.Push(3)
	f.Push(7)
	f.Push(-2)
	f.Push(12)

	var x int = f.Pop()

	if x != 3 || f.last.v != 7 || f.last.prev.v != -2 || f.element.v != 12 {
		t.Fail()
	}
}

func TestFilePop2(t *testing.T) {
	var f File = GetEmptyFile()

	f.Push(32)
	f.Pop()
	if f.last.prev != nil || f.last.next != nil || f.element.prev != nil || f.element.next != nil {
		t.Fail()
	}
}