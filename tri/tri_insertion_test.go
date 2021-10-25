package algolp

import (
	"testing"
)

func TestInsertionVide(t *testing.T){
	var liste []int
	liste = TriInsertion(liste)
	if len(liste) != 0 || liste != nil{
		t.Fail()
	}
}

func TestInsertionVide2(t *testing.T){
	var liste []int = nil
	liste = TriInsertion(liste)
	if len(liste) != 0 || liste != nil {
		t.Fail()
	}
}

func TestDejaTrie(t *testing.T) {
	var liste []int = []int{1, 2, 36, 72, 80}
	liste = TriInsertion(liste)
	if !listeEgal(liste, []int{1, 2, 36, 72, 80}) {
		t.Fail()
	}
}

func TestValeursNegatives(t *testing.T) {
	var liste []int = []int{-5, -50, -63, -05, -98}
	liste = TriInsertion(liste)
	if !listeEgal(liste, []int{-98, -63, -50, -5, -5}) {
		t.Fail()
	}
}


func TestUneValeur(t *testing.T) {
	var liste []int = []int{1}
	liste = TriInsertion(liste)
	if !listeEgal(liste, []int{1}) {
		t.Fail()
	}
}

func TestPlusieursValeurs(t *testing.T) {
	var liste []int = []int{9, 5, 69, 1, 0}
	liste = TriInsertion(liste)
	if !listeEgal(liste, []int{0, 1, 5, 9, 69}) {
		t.Fail()
	}
}

func TestDoublons(t *testing.T) {
	var liste []int = []int{-5, 2, 30, 2, 99, 85, 99, 30, 85, 85, 10}
	liste = TriInsertion(liste)
	if !listeEgal(liste, []int{-5, 2, 2, 10, 30, 30, 85, 85, 85, 99, 99}) {
		t.Fail()
	}
}
func listeEgal(liste1 []int, liste2 []int) (bool) {
	if len(liste1) != len(liste2) {
		return false
	}
	for i := 0; i < len(liste1); i++ {
		if liste1[i] != liste2[i] {
			return false
		}
	}

	return true
}