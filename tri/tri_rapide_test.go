package algolp

import (
	"testing"
)

func TestRapideVide(t *testing.T){
	var liste []int
	TriRapide(liste)
	if len(liste) != 0 || liste != nil{
		t.Fail()
	}
}

func TestRapideVide2(t *testing.T){
	var liste []int = nil
	TriRapide(liste)
	if len(liste) != 0 || liste != nil {
		t.Fail()
	}
}

func TestRapideDejaTrie(t *testing.T) {
	var liste []int = []int{1, 2, 36, 72, 80}
	TriRapide(liste)
	if !listeEgal(liste, []int{1, 2, 36, 72, 80}) {
		t.Fail()
	}
}

func TestRapideValeursNegatives(t *testing.T) {
	var liste []int = []int{-5, -50, -63, -05, -98}
	TriRapide(liste)
	if !listeEgal(liste, []int{-98, -63, -50, -5, -5}) {
		t.Fail()
	}
}


func TestRapideUneValeur(t *testing.T) {
	var liste []int = []int{1}
	TriRapide(liste)
	if !listeEgal(liste, []int{1}) {
		t.Fail()
	}
}

func TestRapidePlusieursValeurs(t *testing.T) {
	var liste []int = []int{9, 5, 69, 1, 0}
	TriRapide(liste)
	if !listeEgal(liste, []int{0, 1, 5, 9, 69}) {
		t.Fail()
	}
}

func TestRapideDoublons(t *testing.T) {
	var liste []int = []int{-5, 2, 30, 2, 99, 85, 99, 30, 85, 85, 10}
	TriRapide(liste)
	if !listeEgal(liste, []int{-5, 2, 2, 10, 30, 30, 85, 85, 85, 99, 99}) {
		t.Fail()
	}
}