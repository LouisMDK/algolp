package algolp

import (
	"testing"
)

func TestSelectionVide(t *testing.T){
	s := Selection([]int{}, 2)
	if s != nil || len(s) != 0{
		t.Fail()
	}
}

func TestSelectionUnElement(t *testing.T){
	s := Selection([]int{10}, 11)
	if !verifierSelection(s, 11){
		t.Fail()
	}
}

func TestSelectionPlusieursElements(t *testing.T){
	s := Selection([]int{3, 4, 5, 12, 24, 8, 9}, 11)
	if !verifierSelection(s, 11){
		t.Fail()
	}
}

func verifierSelection(tab []int, val int) (bool){
	for _, nombre := range tab{
		if nombre > val{
			return false
		}
	}
	return true
}